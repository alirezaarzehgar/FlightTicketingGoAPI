package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
	"github.com/BaseMax/FlightTicketingGoAPI/models"
)

var EXPTIME = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

func HashPassword(pass string) string {
	hashByte := sha256.Sum256([]byte(pass))
	hashStr := hex.EncodeToString(hashByte[:])
	return hashStr
}

func CreateJwtToken(id uint, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(id),
		Issuer:    email,
		ExpiresAt: EXPTIME,
	})
	bearer, _ := token.SignedString(config.GetJwtSecret())
	return bearer
}

func ErrGormToHttp(r *gorm.DB) *echo.HTTPError {
	err := r.Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return echo.ErrNotFound
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return echo.ErrConflict
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return echo.ErrConflict
	case errors.Is(err, gorm.ErrInvalidField):
		return echo.ErrBadRequest
	case err != nil:
		return echo.ErrInternalServerError
	case r.RowsAffected == 0:
		return echo.ErrNotFound
	}
	return nil
}

func Loggedin(c echo.Context) *models.User {
	bearer := c.Request().Header.Get("Authorization")
	token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	email := claims["iss"].(string)
	id, _ := strconv.Atoi(claims["jti"].(string))
	return &models.User{ID: uint(id), Email: email}
}

func IsAdmin(c echo.Context) bool {
	return Loggedin(c).Email == config.GetAdminConf().Email
}

func IsEmployee(c echo.Context) bool {
	return Loggedin(c).Role == models.USERS_ROLE_EMPLOYEE
}

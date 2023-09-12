package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
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
	case errors.Is(err, gorm.ErrForeignKeyViolated):
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return echo.ErrConflict
	case err != nil:
		return echo.ErrInternalServerError
	case r.RowsAffected == 0:
		return echo.ErrNotFound
	}
	return nil
}

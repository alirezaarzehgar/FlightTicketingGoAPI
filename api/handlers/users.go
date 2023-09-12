package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func Register(c echo.Context) error {
	user := models.User{Role: models.USERS_ROLE_PASSENGER}
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = utils.HashPassword(user.Password)
	if err := utils.ErrGormToHttp(db.Create(&user)); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]any{"token": utils.CreateJwtToken(user.ID, user.Email)})
}

func Login(c echo.Context) error {
	var loggedin int64
	var user models.User

	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	fillteredUser := models.User{Email: user.Email, Password: utils.HashPassword(user.Password)}

	db.Where(fillteredUser).First(&models.User{}).Count(&loggedin)

	if loggedin == 0 {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]any{"token": utils.CreateJwtToken(user.ID, user.Email)})
}

func FetchUser(c echo.Context) error {
	return FetchModelById[models.User](c, "id", "password")
}

func FetchUsers(c echo.Context) error {
	return FetchAllModels[models.User](c, "password")
}

func EditUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = utils.HashPassword(user.Password)
	r := db.Model(models.User{}).Where(id).Omit("id, role").Updates(&user)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) error {
	return DeleteById(c, models.User{}, "id")
}

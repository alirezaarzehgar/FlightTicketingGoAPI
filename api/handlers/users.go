package handlers

import (
	"encoding/json"
	"net/http"

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
	return nil
}

func FetchUser(c echo.Context) error {
	return nil
}

func FetchUsers(c echo.Context) error {
	return nil
}

func EditUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func Register(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = utils.HashPassword(user.Password)
	user.Role = models.USERS_ROLE_PASSENGER
	err := utils.ErrGormToHttp(models.Create(&user))
	if err != nil {
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

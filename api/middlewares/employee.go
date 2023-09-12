package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func EmployeePrivilege(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if utils.IsEmployee(c) || utils.IsAdmin(c) {
			return next(c)
		}
		return echo.ErrUnauthorized
	})
}

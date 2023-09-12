package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if utils.IsAdmin(c) {
			return next(c)
		}
		return echo.ErrUnauthorized
	})
}

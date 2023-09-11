package middlewares

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
)

func ValidVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		version, err := strconv.ParseFloat(c.Param("version"), 64)
		if err != nil {
			return echo.ErrNotFound
		}

		for _, v := range config.Versions {
			if v == version {
				return next(c)
			}
		}

		return echo.ErrNotFound
	})
}

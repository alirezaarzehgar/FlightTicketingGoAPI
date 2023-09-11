package middlewares

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ValidIdParams(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if !strings.Contains(c.Path(), "id") {
			return next(c)
		}
		for _, name := range c.ParamNames() {
			if _, err := strconv.ParseUint(c.Param(name), 10, 64); err != nil {
				return echo.ErrNotFound
			}
		}
		return next(c)
	})
}

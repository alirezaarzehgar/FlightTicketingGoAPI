package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func SearchAirline(c echo.Context) error {
	return SearchModel[models.Airline](c)
}

func FetchAllAirlines(c echo.Context) error {
	return FetchAllModels[models.Airline](c, "")
}

func airlineActivity(c echo.Context, isActive bool) error {
	id, _ := strconv.Atoi(c.Param("id"))
	r := db.Model(&models.Airline{}).Where(id).Update("active", isActive)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func ActiveAirline(c echo.Context) error {
	return airlineActivity(c, true)
}

func DeactiveAirline(c echo.Context) error {
	return airlineActivity(c, false)
}

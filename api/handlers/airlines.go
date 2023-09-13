package handlers

import (
	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/labstack/echo/v4"
)

func FetchAirline(c echo.Context) error {
	return FetchModelById[models.Airline](c, "id", "")
}

func FetchAllAirlines(c echo.Context) error {
	return FetchAllModels[models.Airline](c, "")
}

func ActiveAirline(c echo.Context) error {
	return nil
}

func DeactiveAirline(c echo.Context) error {
	return nil
}

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/labstack/echo/v4"
)

const (
	LEAST_FLIGHT_DURATION = time.Minute * 20
)

func NewFlight(c echo.Context) error {
	var flight models.Flight
	if err := json.NewDecoder(c.Request().Body).Decode(&flight); err != nil {
		return echo.ErrBadRequest
	}
	if flight.ArrivalDate.Sub(flight.DepartureDate) <= LEAST_FLIGHT_DURATION {
		return echo.ErrNotAcceptable
	}

	return c.JSON(http.StatusOK, flight)
}

func FetchFlight(c echo.Context) error {
	return nil
}

func FetchAllFlights(c echo.Context) error {
	return nil
}

func EditFlight(c echo.Context) error {
	return nil
}

func DeleteFlight(c echo.Context) error {
	return nil
}

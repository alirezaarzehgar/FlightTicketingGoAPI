package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
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

	r := db.Create(&flight)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	r = db.Preload("Origin").Preload("Destination").First(&flight, flight.ID)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, flight)
}

func SearchFlight(c echo.Context) error {
	return SearchModel[models.Flight](c)
}

func FetchAllFlights(c echo.Context) error {
	return FetchAllModels[models.Flight](c, "")
}

func EditFlight(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var flight models.Flight
	if err := json.NewDecoder(c.Request().Body).Decode(&flight); err != nil {
		return echo.ErrBadRequest
	}
	r := db.Where(id).Updates(&flight)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func DeleteFlight(c echo.Context) error {
	return DeleteById(c, &models.Flight{}, "id")
}

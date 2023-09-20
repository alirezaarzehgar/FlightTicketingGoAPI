package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func passengerNickGen(passengers *[]models.Passenger) *[]models.Passenger {
	var out []models.Passenger
	var emails []string

	for _, passenger := range *passengers {
		emails = append(emails, passenger.Email)
	}

	var registereds []models.Passenger
	db.Find(&registereds, "email IN (?)", emails)

	for _, passenger := range *passengers {
		passenger.Nickname = fmt.Sprintf("%s@%s", passenger.FirstName, passenger.LastName)
		out = append(out, passenger)
		emails = append(emails, passenger.Email)
	}

	for i := 0; i < len(out); i++ {
		for j := 0; j < len(registereds); j++ {
			if out[i].Email == registereds[j].Email {
				out[i].ID = registereds[j].ID
			}
		}
	}

	return &out
}

func Booking(c echo.Context) error {
	ticket := models.Ticket{}
	flightId, _ := strconv.Atoi(c.Param("flight_id"))

	if err := json.NewDecoder(c.Request().Body).Decode(&ticket); err != nil {
		return echo.ErrBadRequest
	}
	ticket.UserID = utils.Loggedin(c).ID
	ticket.FlightID = uint(flightId)
	ticket.BookingDate = time.Now()
	ticket.Passengers = passengerNickGen(ticket.Passengers)

	r := db.Create(&ticket)
	if errors.Is(r.Error, gorm.ErrForeignKeyViolated) {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "flight not found"})
	}
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	db.Preload(clause.Associations).First(&ticket)

	utils.NewTicketSchedule(ticket.ID, ticket.BookingDate)
	return c.JSON(http.StatusOK, ticket)
}

func SearchTicket(c echo.Context) error {
	return SearchModel[models.Ticket](c)
}

func FetchTicket(c echo.Context) error {
	var ticket models.Ticket
	id, _ := strconv.Atoi(c.Param("id"))
	r := db.Preload(clause.Associations).First(&ticket, id)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	db.Preload(clause.Associations).First(&ticket.Flight)
	return c.JSON(http.StatusOK, ticket)
}

func FetchAllTicketsByFlight(c echo.Context) error {
	flightId, _ := strconv.Atoi(c.Param("id"))

	var ticket []models.Ticket
	r := db.Scopes(Paginate(c)).Preload(clause.Associations).Find(&ticket, "flight_id = ?", flightId)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ticket)
}

func EditTicket(c echo.Context) error {
	ticket := models.Ticket{}
	ticketId, _ := strconv.Atoi(c.Param("id"))

	if err := json.NewDecoder(c.Request().Body).Decode(&ticket); err != nil {
		return echo.ErrBadRequest
	}
	ticket.BookingDate = time.Now()
	ticket.Passengers = passengerNickGen(ticket.Passengers)

	r := db.Where(ticketId).Updates(ticket)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func DeleteTicket(c echo.Context) error {
	ticketId, _ := strconv.Atoi(c.Param("id"))
	err := DeleteById(c, models.Ticket{}, "id")
	utils.CancelTicket(uint(ticketId))
	return err
}

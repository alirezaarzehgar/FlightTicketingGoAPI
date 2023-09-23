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
	db.Preload(clause.Associations).First(&ticket.User)
	db.Preload(clause.Associations).First(&ticket.Flight)

	utils.NewTicketSchedule(ticket)

	strPassengers := ""
	for _, passenger := range *ticket.Passengers {
		strPassengers += fmt.Sprintf("%s %s (%s).\n", passenger.FirstName, passenger.LastName, passenger.Email)
	}

	sub := "Flight Booking"
	body := "Dear passenger,\n\n" +
		"Please pay ticket passage and then you can flight.\n" +
		"If you don't pay this ticket passage your ticket will canceled after 10 minutes.\n" +
		"Ticket details:\n" +
		"Origin: " + ticket.Flight.Origin.Name + " on " + ticket.Flight.Origin.Country + ".\n" +
		"Destination: " + ticket.Flight.Destination.Name + " on " + ticket.Flight.Destination.Country + ".\n" +
		"Flight start-end clock: " + ticket.Flight.DepartureDate.String() + "-" + ticket.Flight.ArrivalDate.String() + ".\n\n" +
		"Passengers: \n" + strPassengers +
		"Total Price: " + fmt.Sprint(ticket.TotalPrice) + ".\n" +
		"Payment gateway: " + PAYMENT_LINK + "\n"

	go utils.EasySendMail(sub, body, ticket.User.Email)

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

func FetchAllTickets(c echo.Context) error {
	var tickets []models.Ticket
	r := db.Scopes(Paginate(c)).Preload(clause.Associations).Find(&tickets, "user_id = ?", utils.Loggedin(c).ID)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tickets)
}

func FetchAllTicketsByFlight(c echo.Context) error {
	flightId, _ := strconv.Atoi(c.Param("id"))

	var tickets []models.Ticket
	r := db.Scopes(Paginate(c)).Preload(clause.Associations).Find(&tickets, "flight_id = ?", flightId)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tickets)
}

func EditTicket(c echo.Context) error {
	ticketId, _ := strconv.Atoi(c.Param("id"))
	ticket := models.Ticket{ID: uint(ticketId)}

	if !utils.IsTicketOnTime(uint(ticketId)) {
		return echo.ErrNotFound
	}

	if err := json.NewDecoder(c.Request().Body).Decode(&ticket); err != nil {
		return echo.ErrBadRequest
	}
	ticket.BookingDate = time.Now()
	passengers := passengerNickGen(ticket.Passengers)

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(passengers)
	db.Model(&ticket).Association("Passengers").Clear()
	ticket.Passengers = passengers

	r := db.Updates(ticket)
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

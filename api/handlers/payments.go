package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/payment"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func CreatePaymentTransaction(c echo.Context) error {
	ticketId, _ := strconv.Atoi(c.Param("ticket_id"))
	ticket := models.Ticket{}

	r := db.First(&ticket, ticketId)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	var transRepteat int64
	db.Where(&models.Transaction{TicketID: ticket.ID, Success: true}).
		Find(&models.Transaction{}).Count(&transRepteat)
	if transRepteat > 0 {
		return c.JSON(http.StatusConflict, map[string]any{"message": "Your ticket has already been paid"})
	}

	trans := models.Transaction{
		TicketID: ticket.ID,
		Amount:   uint(ticket.TotalPrice),
	}
	r = db.Create(&trans)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}

	gw := payment.NewAqayePardakht("sandbox", utils.GetRepeatedUrl(c))
	authority, err := gw.Request(uint(ticket.TotalPrice), trans.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	trans.Authority = authority
	db.Save(trans)

	url := gw.CreateRequestUrl(authority)
	return c.JSON(http.StatusOK, map[string]any{"url": url})
}

func SuccessTransaction(c echo.Context) error {
	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))
	trans := models.Transaction{ID: uint(transactionId), Success: true}
	r := db.Clauses(clause.Returning{}).Updates(&trans)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, trans)
}

func VerfifyPayment(c echo.Context) error {
	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))
	trans := models.Transaction{}
	r := db.First(&trans, transactionId)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	gw := payment.NewAqayePardakht("sandbox", utils.GetRepeatedUrl(c))
	verify, err := gw.Veify(trans.Amount, trans.Authority)
	if err != nil {
		return echo.ErrBadRequest
	}
	status := "failed"
	if verify {
		status = "success"
	}
	return c.JSON(http.StatusOK, map[string]any{"status": status})
}

func SearchPayments(c echo.Context) error {
	return SearchModels[models.Transaction](c)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

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

	gw := payment.NewAqayePardakht("sandbox", utils.GetRepeatedUrl(c))
	url, err := payment.CreateRequest(gw, uint(ticket.TotalPrice))
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]any{"url": url})
}

func DoneTransaction(c echo.Context) error {
	return nil
}

func VerfifyPayment(c echo.Context) error {
	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))
	trans := models.Transaction{}
	r := db.First(&trans, transactionId)
	if err := utils.ErrGormToHttp(r); err != nil {
		return err
	}
	gw := payment.NewAqayePardakht("sandbox", utils.GetRepeatedUrl(c))
	payment.Verify(gw, trans.Amount, trans.Authority)
	return nil
}

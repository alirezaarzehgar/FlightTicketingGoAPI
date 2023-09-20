package utils

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
)

var (
	ticketIds             = make(map[uint]time.Time)
	expireTime            = time.Minute * 10
	removeWorkerCycleTime = time.Second * 5
	db                    *gorm.DB
)

func initStoredTickets() {
	var tickets []models.Ticket
	db.Find(&tickets, "paid = ?", false)

	for _, ticket := range tickets {
		NewTicketSchedule(ticket.ID, ticket.BookingDate)
	}
}

func cancelExpiredTickets() {
	for {
		for ticketId := range ticketIds {
			if !IsTicketOnTime(ticketId) {
				CancelTicket(ticketId)
			}
		}
		time.Sleep(removeWorkerCycleTime)
	}
}

func RunTicketWorkers(externalDB *gorm.DB) {
	db = externalDB

	initStoredTickets()
	go cancelExpiredTickets()
}

func NewTicketSchedule(ticketId uint, time time.Time) {
	ticketIds[ticketId] = time
}

func IsTicketOnTime(ticketId uint) bool {
	return time.Since(ticketIds[ticketId]) <= expireTime
}

func CancelTicket(ticketId uint) {
	var ticket models.Ticket
	db.Where("paid = ?", false).First(&ticket, ticketId)
	db.Model(&ticket).Association("Passengers").Clear()
	db.Select(clause.Associations).Delete(ticket)
	delete(ticketIds, ticketId)
}

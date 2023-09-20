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
		NewTicketSchedule(ticket.ID)
	}
}

func cancelExpiredTickets() {
	for {
		for ticketId := range ticketIds {
			if !IsTicketOnTime(ticketId) {
				db.Select(clause.Associations).Where("paid = ?", false).Delete(models.Ticket{}, ticketId)
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

func NewTicketSchedule(ticketId uint) {
	ticketIds[ticketId] = time.Now()
}

func IsTicketOnTime(ticketId uint) bool {
	return time.Since(ticketIds[ticketId]) <= expireTime
}

func CancelTicket(ticketId uint) {
	ticketIds[ticketId] = time.Time{}
}

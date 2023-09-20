package utils

import (
	"time"
)

var ticketIds map[uint]time.Time

func NewTicketSchedule(ticketId uint) {
	ticketIds[ticketId] = time.Now()
}

func IsTicketOnTime(ticketId uint) bool {
	return time.Since(ticketIds[ticketId]) <= time.Minute*10
}

func CancelTicketTimer(ticketId uint) {
	ticketIds[ticketId] = time.Time{}
}

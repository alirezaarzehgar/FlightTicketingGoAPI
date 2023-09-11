package models

import "time"

type Ticket struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	FlightID    uint      `gorm:"not null" json:"flight_id"`
	Flight      Flight    `json:"flight,omitempty"`
	Users       []User    `gorm:"many2many:user_tickets" json:"passengers"`
	TotalPrice  float64   `gorm:"-" json:"total_price"`
	BookingDate time.Time `gorm:"not null" json:"booking_date"`
}

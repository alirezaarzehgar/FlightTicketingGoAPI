package models

import "time"

type Passenger struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	TicketID  uint   `json:"ticket_id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null; unique" json:"email"`
}

type Ticket struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `gorm:"not null" json:"user_id"`
	FlightID    uint        `gorm:"not null" json:"flight_id"`
	Flight      Flight      `json:"flight,omitempty"`
	Passengers  []Passenger `json:"passengers"`
	TotalPrice  float64     `gorm:"-" json:"total_price"`
	BookingDate time.Time   `gorm:"not null" json:"booking_date"`
}

package models

import "time"

const (
	USERS_ROLE_ADMIN     = "admin"
	USERS_ROLE_EMPLOYEE  = "employee"
	USERS_ROLE_PASSENGER = "passenger"
)

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Email    string   `gorm:"not null;unique" json:"email,omitempty"`
	Password string   `gorm:"not null" json:"password,omitempty"`
	Role     string   `gorm:"not null" json:"role,omitempty"`
	Tickets  []Ticket `gorm:"many2many:user_tickets" json:"tickets,omitempty"`
}

type Airline struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	Callsign string `json:"callsign"`
	Country  string `gorm:"not null" json:"country"`
	Active   bool   `gorm:"not null" json:"active"`
}

type Flight struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OriginID      uint      `gorm:"not null" json:"origin_id"`
	Origin        Airline   `gorm:"foreignKey:OriginID" json:"origin,omitempty"`
	DestinationID uint      `gorm:"not null" json:"destination_id"`
	Destination   Airline   `gorm:"foreignKey:DestinationID" json:"destination,omitempty"`
	DepartureDate time.Time `gorm:"not null" json:"departure_date"`
	ArrivalDate   time.Time `gorm:"not null" json:"arrival_date"`
	Price         float64   `gorm:"not null" json:"price"`
}

type Passenger struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	TicketID  uint   `json:"ticket_id"`
	Ticket    Ticket `gorm:"foreignKey:ID"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null; unique" json:"email"`
}

type Ticket struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	UserID      uint        `gorm:"foreignKey:ID;not null" json:"user_id"`
	FlightID    uint        `gorm:"not null" json:"flight_id"`
	Flight      Flight      `gorm:"foreignKey:ID" json:"flight,omitempty"`
	Passengers  []Passenger `gorm:"foreignKey:ID" json:"passengers"`
	TotalPrice  float64     `gorm:"-" json:"total_price"`
	BookingDate time.Time   `gorm:"not null" json:"booking_date"`
}

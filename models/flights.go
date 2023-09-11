package models

import (
	"time"
)

type Flight struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Origin        string    `gorm:"not null" json:"origin"`
	Destination   string    `gorm:"not null" json:"destination"`
	DepartureDate time.Time `gorm:"not null" json:"departure_date"`
	ArrivalDate   time.Time `gorm:"not null" json:"arrival_date"`
	Price         float64   `gorm:"not null" json:"price"`
}

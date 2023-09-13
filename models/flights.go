package models

import (
	"time"
)

type Flight struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OriginID      uint      `gorm:"not null" json:"origin"`
	Origin        Airline   `gorm:"foreignKey:OriginID" json:"origin_id"`
	DestinationID uint      `gorm:"not null" json:"destination_id"`
	Destination   Airline   `gorm:"foreignKey:DestinationID" json:"destination"`
	DepartureDate time.Time `gorm:"not null" json:"departure_date"`
	ArrivalDate   time.Time `gorm:"not null" json:"arrival_date"`
	Price         float64   `gorm:"not null" json:"price"`
}

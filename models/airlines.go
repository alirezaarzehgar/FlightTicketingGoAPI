package models

type Airline struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	Callsign string `json:"callsign"`
	Country  string `gorm:"not null" json:"country"`
	Active   bool   `gorm:"not null" json:"active"`
}

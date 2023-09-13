package database

import (
	"encoding/json"
	"os"

	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
	"github.com/BaseMax/FlightTicketingGoAPI/models"
)

func AirlinesMigrate(db *gorm.DB) error {
	var airlines []models.Airline
	data, err := os.ReadFile(config.GetAirlineConfPath())
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &airlines); err != nil {
		return err
	}

	return db.CreateInBatches(airlines, len(airlines)).Error
}

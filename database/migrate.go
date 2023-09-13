package database

import (
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/models"
)

func Migrate(db *gorm.DB) error {
	if db.Migrator().HasTable(&models.User{}) {
		return nil
	}

	err := db.AutoMigrate(&models.User{}, &models.Passenger{}, &models.Flight{}, &models.Ticket{}, &models.Airline{})
	if err != nil {
		return err
	}

	if err := UsersMigrate(db); err != nil {
		return err
	}

	if err := AirlinesMigrate(db); err != nil {
		return err
	}

	return nil
}

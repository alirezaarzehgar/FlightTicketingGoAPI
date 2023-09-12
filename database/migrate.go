package database

import (
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func Migrate(db *gorm.DB) error {
	if db.Migrator().HasTable(&models.User{}) {
		return nil
	}

	err := db.AutoMigrate(&models.User{}, &models.Passenger{}, &models.Flight{}, &models.Ticket{})
	if err != nil {
		return err
	}

	conf := config.GetAdminConf()
	admin := models.User{
		Email:    conf.Email,
		Password: utils.HashPassword(conf.Password),
		Role:     models.USERS_ROLE_ADMIN,
	}
	return db.Create(&admin).Error
}

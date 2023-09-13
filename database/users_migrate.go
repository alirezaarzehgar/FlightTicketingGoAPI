package database

import (
	"gorm.io/gorm"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
	"github.com/BaseMax/FlightTicketingGoAPI/models"
	"github.com/BaseMax/FlightTicketingGoAPI/utils"
)

func UsersMigrate(db *gorm.DB) error {
	conf := config.GetAdminConf()
	admin := models.User{
		Email:    conf.Email,
		Password: utils.HashPassword(conf.Password),
		Role:     models.USERS_ROLE_ADMIN,
	}
	return db.Create(&admin).Error
}

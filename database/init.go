package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
)

func InitDB(c *config.DbConf) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		c.Host, c.User, c.Password, c.DbName, c.Port, c.TimeZone)

	config := &gorm.Config{TranslateError: true}
	if !c.Debug {
		config.Logger = logger.Discard
	}

	db, err = gorm.Open(postgres.Open(dsn), config)
	return
}

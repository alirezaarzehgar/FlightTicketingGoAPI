package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/BaseMax/FlightTicketingGoAPI/api/routes"
	"github.com/BaseMax/FlightTicketingGoAPI/config"
	"github.com/BaseMax/FlightTicketingGoAPI/database"
	"github.com/BaseMax/FlightTicketingGoAPI/models"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("faild to load .env: ", err)
	}

	dbConf, err := config.GetDbConfig()
	if err != nil {
		log.Fatal(".env: ", err)
	}

	db, err := database.InitDB(dbConf)
	if err != nil {
		log.Fatal("db: ", err)
	}
	models.SetDB(db)

	r := routes.InitRoutes()
	httpConf := config.GetHttpConfig()
	r.Logger.Fatal(r.Start(httpConf.Addr))
}

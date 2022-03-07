// models/setup.go

package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	if err := database.AutoMigrate(&Users{}, &Tags{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}
	if err := database.AutoMigrate(&Questions{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}
	if err := database.AutoMigrate(&Answers{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}

	DB = database
}

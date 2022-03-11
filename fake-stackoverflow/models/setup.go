// models/setup.go

package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	connection_string := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	database, err := gorm.Open(postgres.Open(connection_string))

	if err != nil {
		panic("Failed to connect to database!")
	}

	if err := database.AutoMigrate(&Users{}, &Tag{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}
	if err := database.AutoMigrate(&Question{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}
	if err := database.AutoMigrate(&Answer{}); err != nil {
		log.Print("Error AutoMigrate: ", err)
	}

	DB = database
}

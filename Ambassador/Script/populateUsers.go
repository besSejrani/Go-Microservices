package main

import (
	config "ambassador/Config"
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/bxcodec/faker/v3"
)

func main() {

	// Environment variables
	config.Secrets()

	// Database
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassadors := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassadors.SetPassword("123456789")

		database.DB.Create(&ambassadors)
	}
}

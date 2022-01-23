package main

import (
	config "ambassador/Config"
	database "ambassador/Database"
	models "ambassador/Models"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {

	// Environment variables
	config.Secrets()

	// Database
	database.Connect()

	for i := 0; i < 30; i++ {
		products := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&products)
	}
}

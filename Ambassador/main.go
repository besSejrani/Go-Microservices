package main

import (
	config "ambassador/Config"
	database "ambassador/Database"
	routes "ambassador/Routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Environment variables
	config.Secrets()

	// Database
	database.Connect()

	app := fiber.New()

	// CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Routes
	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

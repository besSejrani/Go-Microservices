package routes

import (
	"ambassador/Controllers/ambassador"
	"ambassador/Controllers/auth"
	"ambassador/Controllers/product"
	user "ambassador/Controllers/user"
	middlewares "ambassador/Middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Prefix
	api := app.Group("api")
	admin := api.Group("admin")

	//-----------------------------------------
	// Public Routes
	//-----------------------------------------

	// Login & Register Routes
	admin.Post("/register", auth.Register)
	admin.Post("/login", auth.Login)

	//-----------------------------------------
	// Private Routes
	//-----------------------------------------

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Post("/logout", auth.Logout)
	adminAuthenticated.Get("/user", user.GetUser)

	// Users Routes
	adminAuthenticated.Put("/users/info", user.UpdateUser)
	adminAuthenticated.Put("/users/password", user.UpdatePassword)

	// Ambassadors Routes
	adminAuthenticated.Get("/ambassadors", ambassador.GetAmbassadors)

	// Products Routes
	adminAuthenticated.Get("/product/:id", product.GetProduct)
	adminAuthenticated.Get("/products", product.GetProducts)
	adminAuthenticated.Post("/product/create", product.CreateProduct)
	adminAuthenticated.Put("/product/:id", product.UpdateProduct)
	adminAuthenticated.Delete("/product/:id", product.DeleteProduct)

}

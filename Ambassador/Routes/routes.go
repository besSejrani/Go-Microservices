package routes

import (
	"ambassador/Controllers/ambassador"
	"ambassador/Controllers/auth"
	"ambassador/Controllers/product"
	user "ambassador/Controllers/user"
	middlewares "ambassador/Middlewares"
	"html/template"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// app.Get("/*", swagger.New(swagger.ConfigDefault)) // default

	app.Get("/*", swagger.New(swagger.Config{ // custom

		Title:  "Swagger UI",
		Layout: "StandaloneLayout",
		Plugins: []template.JS{
			template.JS("SwaggerUIBundle.plugins.DownloadUrl"),
		},
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
		DeepLinking:              true,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DefaultModelRendering:    "example",
		DocExpansion:             "list",
		ShowMutatedRequest:       true,
	}))

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

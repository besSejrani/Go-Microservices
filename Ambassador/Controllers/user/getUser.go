package auth

import (
	database "ambassador/Database"
	middlewares "ambassador/Middlewares"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	id, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)

}

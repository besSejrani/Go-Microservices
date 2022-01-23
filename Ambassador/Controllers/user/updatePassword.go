package auth

import (
	database "ambassador/Database"
	middlewares "ambassador/Middlewares"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func UpdatePassword(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{
		Id: id,
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

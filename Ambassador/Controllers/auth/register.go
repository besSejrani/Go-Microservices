package auth

import (
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		FirstName:    data["firstName"],
		LastName:     data["lastName"],
		Email:        data["email"],
		IsAmbassador: false,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

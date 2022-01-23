package product

import (
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

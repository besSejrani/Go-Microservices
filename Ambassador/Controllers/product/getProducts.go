package product

import (
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {

	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)

}

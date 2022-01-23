package product

import (
	database "ambassador/Database"
	models "ambassador/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {

	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.ID = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

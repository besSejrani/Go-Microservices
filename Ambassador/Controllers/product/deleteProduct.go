package product

import (
	database "ambassador/Database"
	models "ambassador/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.ID = uint(id)

	database.DB.Delete(&product)

	return nil
}

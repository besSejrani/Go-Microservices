package product

import (
	database "ambassador/Database"
	models "ambassador/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateProduct(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(data); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		ID: uint(id),
	}

	database.DB.Model(&product).Updates(&product)

	return c.JSON(product)

}

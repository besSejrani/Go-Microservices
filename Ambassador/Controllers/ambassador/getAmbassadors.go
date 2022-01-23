package ambassador

import (
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

func GetAmbassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(&users)
}

package ambassador

import (
	database "ambassador/Database"
	models "ambassador/Models"

	"github.com/gofiber/fiber/v2"
)

// GetBooks func gets all exists books.
// @Summary      get all ambassadors
// @Description  Get all ambassadors.
// @Tags         Ambassadors
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /api/admin/ambassadors [get]
func GetAmbassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(&users)
}

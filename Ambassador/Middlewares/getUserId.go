package middlewares

import (
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}

package middlewares

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id uint) (string, error) {
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	// payload.Scope = scope

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

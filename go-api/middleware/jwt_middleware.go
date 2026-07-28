package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("matrix-app-secret")

func JWTProtected(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"error": "missing authorization header",
			})
	}

	tokenString := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"error": "invalid token",
			})
	}

	return c.Next()
}
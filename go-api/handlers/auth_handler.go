package handlers

import (
	"matrix-app/go-api/models"
	"matrix-app/go-api/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var request models.LoginRequest

	if err := c.BodyParser(&request); err != nil {

		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": "invalid request",
			})
	}

	if request.Username != "admin" ||
		request.Password != "admin123" {

		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"error": "invalid credentials",
			})
	}

	token, err := services.GenerateToken(
		request.Username,
	)

	if err != nil {

		return c.Status(
			fiber.StatusInternalServerError,
		).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: token,
	})
}
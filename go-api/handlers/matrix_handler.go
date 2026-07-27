package handlers

import (
	"matrix-app/go-api/models"
	"matrix-app/go-api/services"

	"github.com/gofiber/fiber/v2"
)

func ValidateMatrix(c *fiber.Ctx) error {

	var request models.MatrixRequest

	if err := c.BodyParser(&request); err != nil {

    return c.Status(fiber.StatusBadRequest).
        JSON(fiber.Map{
            "error": "invalid request body",
        })
}

println("Rows:", len(request.Matrix))

	if err := services.ValidateMatrix(
		request.Matrix,
	); err != nil {

		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"valid": true,
	})
}
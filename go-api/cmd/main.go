package main

import (
	"log"

	"matrix-app/go-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/api/qr", handlers.ValidateMatrix)

	log.Fatal(app.Listen(":8080"))
}
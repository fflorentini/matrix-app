package main

import (
	"log"

	"matrix-app/go-api/handlers"
	"matrix-app/go-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/login", handlers.Login)

	app.Post(
		"/api/qr",
		middleware.JWTProtected,
		handlers.ValidateMatrix,
	)

	log.Fatal(app.Listen(":8080"))
}
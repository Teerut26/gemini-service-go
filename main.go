package main

import (
	"gemini-service/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	apiRoute := app.Group("/api")
	v1Route := apiRoute.Group("/v1")
	v1Route.Post("/gemini", services.GeminiHandler)

	app.Listen(":3000")
}

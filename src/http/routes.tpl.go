package http

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes register web routes
func RegisterRoutes(router *fiber.App) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to bopher app")
	})
}

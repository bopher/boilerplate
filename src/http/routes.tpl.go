package http

import (
	"mekramy/__boiler/src/app"
	"time"

	"github.com/bopher/http/middlewares"
	"github.com/bopher/utils"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes register web routes
func RegisterRoutes(router *fiber.App) {
	// Global Middlewares
	router.Use(middlewares.RateLimiter(
		"GLOBAL-LIMITER",
		app.Config().Cast("web.limit").UInt32Safe(60),
		time.Minute,
		app.Cache(), func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		})) // Accept 60 request in minutes
	router.Use(func(c *fiber.Ctx) error {
		if ok := utils.BoolOrPanic(app.IsUnderMaintenance()); ok {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		return c.Next()
	}) // Maintenance mode
	// router.Use(middlewares.CSRFMiddleware(mySession))
	// router.Use(middlewares.JSONOnly(nil))

	// Routes
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to bopher app")
	})

	// Fallback
	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}

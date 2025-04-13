package health

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	// Health check route
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Health check route with a delay
	router.Get("/delay", func(c *fiber.Ctx) error {
		time.Sleep(5 * time.Second) // Simulate a delay of 5 seconds
		return c.SendString("OK")
	})
}

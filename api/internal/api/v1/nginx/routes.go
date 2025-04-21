package nginx

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {

	router.Post("/reload", ReloadNginx)
	router.Get("/status", StatusNginx)
}

func ReloadNginx(c *fiber.Ctx) error {
	cmd := exec.Command("/var/starpanel/packages/nginx/v1/nginx", "-s", "reload")
	err := cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to reload nginx: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Nginx reloaded successfully",
	})
}

func StatusNginx(c *fiber.Ctx) error {
	cmd := exec.Command("nginx", "-t")
	err := cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Nginx is not running: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Nginx is running",
	})
}

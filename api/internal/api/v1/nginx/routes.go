package nginx

import (
	"fmt"
	"os/exec"

	"github.com/arilsonb/starpanel/cmd/queue"
	"github.com/arilsonb/starpanel/internal/api/v1/nginx/tasks"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {

	router.Post("/reload", ReloadNginx)
	router.Get("/status", StatusNginx)
	router.Get("/install", InstallNginx)
}

func InstallNginx(c *fiber.Ctx) error {
	task, err := tasks.InstallNginxTask("nginx", "/var/starpanel/packages/nginx/v1", "1.24.0", "https://nginx.org/download/nginx-1.24.0.tar.gz")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create task: %v", err),
		})
	}
	
	info, err := queue.Client.Enqueue(task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to enqueue task: %v", err),
		})
	}
	fmt.Printf("Enqueued task: id=%s queue=%s  args=%s\n", info.ID, info.Queue, task.Payload())
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Task enqueued: %s", info.ID),
	})
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

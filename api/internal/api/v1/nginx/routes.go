package nginx

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {

	router.Post("/reload", ReloadNginx)
	router.Get("/status", StatusNginx)
	router.Get("/install", InstallNginx)
}

func InstallNginx(c *fiber.Ctx) error {
	cmd := exec.Command("bash", "-c", "wget https://nginx.org/download/nginx-1.24.0.tar.gz")
	err := cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to install nginx: %v", err),
		})
	}

	cmd = exec.Command("bash", "-c", "tar -xzf nginx-1.24.0.tar.gz")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to extract nginx: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "cd nginx-1.24.0 && ./configure && make && make install")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to build nginx: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "rm -rf nginx-1.24.0 nginx-1.24.0.tar.gz")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to clean up nginx files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create nginx directory: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "cp -r /usr/local/nginx /var/starpanel/packages/nginx/v1")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to copy nginx files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "rm -rf /usr/local/nginx")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to remove nginx files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "ln -s /var/starpanel/packages/nginx/v1/nginx /usr/local/bin/nginx")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create nginx symlink: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "ln -s /var/starpanel/packages/nginx/v1/nginx /usr/bin/nginx")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create nginx symlink: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1/conf")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create nginx conf directory: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "cp -r /var/starpanel/packages/nginx/v1/nginx/conf/* /var/starpanel/packages/nginx/v1/conf")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to copy nginx conf files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "rm -rf /var/starpanel/packages/nginx/v1/nginx/conf")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to remove nginx conf files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1/logs")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create nginx logs directory: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "cp -r /var/starpanel/packages/nginx/v1/nginx/logs/* /var/starpanel/packages/nginx/v1/logs")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to copy nginx logs files: %v", err),
		})
	}
	cmd = exec.Command("bash", "-c", "rm -rf /var/starpanel/packages/nginx/v1/nginx/logs")
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to remove nginx logs files: %v", err),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Nginx installed successfully",
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

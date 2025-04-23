package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/hibiken/asynq"
)

func HandleinstallNginxTask(ctx context.Context, t *asynq.Task) error {
	var payload InstallNginxPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	fmt.Printf("Installing Nginx: %s from %s\n", payload.PackageName, payload.PackagePath)

	cmd := exec.Command("bash", "-c", "wget -P ./installers https://nginx.org/download/nginx-1.24.0.tar.gz")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to download nginx: %v", err)
	}

	cmd = exec.Command("bash", "-c", "tar -xzf ./installers/nginx-1.24.0.tar.gz")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to extract nginx: %v", err)
	}
	fmt.Printf("Extracted Nginx: %s\n", payload.PackageName)

	cmd = exec.Command("bash", "-c", "cd ./installers/nginx-1.24.0 && ./configure && make && make install 2>&1 | tee output.log")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install nginx: %v", err)
	}
	fmt.Printf("Installed Nginx: %s\n", payload.PackageName)

	cmd = exec.Command("bash", "-c", "rm -rf ./installers/nginx-1.24.0 ./installers/nginx-1.24.0.tar.gz")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clean up nginx files: %v", err)
	}
	cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create nginx package directory: %v", err)
	}
	cmd = exec.Command("bash", "-c", "cp -r /usr/local/nginx /var/starpanel/packages/nginx/v1")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to copy nginx files: %v", err)
	}
	cmd = exec.Command("bash", "-c", "rm -rf /usr/local/nginx")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to remove nginx files: %v", err)
	}
	cmd = exec.Command("bash", "-c", "ln -s /var/starpanel/packages/nginx/v1/nginx /usr/local/bin/nginx")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create nginx symlink: %v", err)
	}
	cmd = exec.Command("bash", "-c", "ln -s /var/starpanel/packages/nginx/v1/nginx /usr/bin/nginx")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create nginx symlink: %v", err)
	}
	cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1/conf")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create nginx conf directory: %v", err)
	}
	cmd = exec.Command("bash", "-c", "cp -r /var/starpanel/packages/nginx/v1/nginx/conf/* /var/starpanel/packages/nginx/v1/conf")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to copy nginx conf files: %v", err)
	}
	cmd = exec.Command("bash", "-c", "rm -rf /var/starpanel/packages/nginx/v1/nginx/conf")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to remove nginx conf files: %v", err)
	}

	fmt.Printf("Nginx installed successfully: %s\n", payload.PackageName)

	// cmd = exec.Command("bash", "-c", "mkdir -p /var/starpanel/packages/nginx/v1/logs")
	// err = cmd.Run()
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": fmt.Sprintf("Failed to create nginx logs directory: %v", err),
	// 	})
	// }
	// cmd = exec.Command("bash", "-c", "cp -r /var/starpanel/packages/nginx/v1/nginx/logs/* /var/starpanel/packages/nginx/v1/logs")
	// err = cmd.Run()
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": fmt.Sprintf("Failed to copy nginx logs files: %v", err),
	// 	})
	// }
	// cmd = exec.Command("bash", "-c", "rm -rf /var/starpanel/packages/nginx/v1/nginx/logs")
	// err = cmd.Run()
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": fmt.Sprintf("Failed to remove nginx logs files: %v", err),
	// 	})
	// }

	return nil
}

package v1

import (
	"github.com/arilsonb/starpanel/internal/api/v1/nginx"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	nginxGroup := router.Group("/nginx")
	nginx.RegisterRoutes(nginxGroup)
}

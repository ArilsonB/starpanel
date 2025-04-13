package main

import (
	"github.com/arilsonb/starpanel/internal/api/auth"
	"github.com/arilsonb/starpanel/internal/api/health"
	v1 "github.com/arilsonb/starpanel/internal/api/v1"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	authGroup := api.Group("/auth")
	auth.RegisterRoutes(authGroup)

	healthGroup := api.Group("/health")
	health.RegisterRoutes(healthGroup)

	v1Group := api.Group("/v1")
	v1.RegisterRoutes(v1Group)

	app.Static("/", "./public")

	app.Listen(":8080")
}

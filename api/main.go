package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true})
	})

	app.Static("/", "./public")

	app.Listen(":8080")
}


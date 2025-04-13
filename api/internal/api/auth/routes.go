package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	// Define the routes for authentication here
	router.Post("/signin", SignIn)
	router.Post("/signout", SignOut)
}

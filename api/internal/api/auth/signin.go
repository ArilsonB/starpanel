package auth

import "github.com/gofiber/fiber/v2"

var jwtSecret = []byte("secret_key") // Replace with your actual secret key

func JWTSecret() []byte {
	return jwtSecret
}

func SignIn(c *fiber.Ctx) error {
	// Handle sign-in logic here
	return nil
}

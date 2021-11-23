package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("i-token")

	// LOGIC FOR AUTHENTICATION

	if token != "haha" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": "invalid token",
		})
	}
	return ctx.Next()
}

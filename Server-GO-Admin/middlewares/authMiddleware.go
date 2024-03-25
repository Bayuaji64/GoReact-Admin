package middlewares

import (
	"example.com/go-admin/utility"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, claims, err := utility.ParseJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	// Menyimpan claims ke context untuk digunakan di handlers berikutnya
	c.Locals("claims", claims)

	return c.Next()
}

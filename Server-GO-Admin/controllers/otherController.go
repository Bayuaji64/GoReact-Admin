package controllers

import "github.com/gofiber/fiber/v3"

func Other(c fiber.Ctx) error {
	return c.SendString("Yang lain 121111")
}

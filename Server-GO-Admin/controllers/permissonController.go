package controllers

import (
	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Allpermission(c *fiber.Ctx) error {
	var p []models.Permission

	db.DB.Find(&p)

	return c.JSON(p)
}

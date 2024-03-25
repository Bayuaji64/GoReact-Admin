package controllers

import (
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Alluser(c *fiber.Ctx) error {
	var users []models.User

	db.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	// Cek apakah email sudah ada
	var existingUser models.User
	result := db.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Email already exists"})
	}

	user.SetPassword("1234")

	if err := db.DB.Create(&user).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create user"})
	}

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.DB.Find(&user)

	return c.JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	db.DB.Model(&user).Updates(user)

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "Delete user success"})

}

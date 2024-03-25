package controllers

import (
	"time"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"example.com/go-admin/utility"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
	}

	var existingAdmin models.User
	result := db.DB.Where("email = ?", user.Email).First(&existingAdmin)
	if result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Email already exists"})
	}

	user.SetPassword(data["password"])

	db.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// Parse input data
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	// Mencari user berdasarkan email
	result := db.DB.Where("email = ?", data["email"]).First(&user)
	if result.Error != nil || user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Memverifikasi password
	if err := user.ComparePassword(data["password"]); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	tokenString, err := utility.GenerateJWT(int(user.Id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "Login success"})
}

func User(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*utility.Claims) // Pastikan casting sesuai dengan tipe claims Anda

	var user models.User

	result := db.DB.Where("id = ?", claims.Issuer).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout success",
	})

}

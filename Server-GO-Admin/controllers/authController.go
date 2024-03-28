package controllers

import (
	"strconv"
	"time"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"example.com/go-admin/utility"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		RoleId:    1,
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

// func User(c *fiber.Ctx) error {
// 	claims := c.Locals("claims").(*utility.Claims) // Pastikan casting sesuai dengan tipe claims Anda

// 	var user models.User

// 	result := db.DB.Where("id = ?", claims.Issuer).First(&user)
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
// 	}

// 	return c.JSON(user)
// }

func User(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*utility.Claims) // Pastikan casting sesuai dengan tipe claims Anda

	var user models.User

	// Menambahkan Preload("Role") untuk memuat data Role bersamaan dengan User
	result := db.DB.Where("id = ?", claims.Issuer).Preload("Role").First(&user)
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

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	// Parse request body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	// Ambil claims dari JWT atau session (sesuaikan dengan implementasi Anda)
	claims, ok := c.Locals("claims").(*utility.Claims) // Pastikan casting sesuai dengan tipe claims Anda
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	userId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	userUpdates := models.User{
		Firstname: data["first_name"],
		Lastname:  data["last_name"],
		Email:     data["email"],
	}

	result := db.DB.Model(&models.User{}).Where("id = ?", userId).Updates(userUpdates)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user information"})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	var updatedUser models.User
	if err := db.DB.Where("id = ?", userId).First(&updatedUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to retrieve updated user information"})
	}

	return c.JSON(updatedUser)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	claims := c.Locals("claims").(*utility.Claims)

	userId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	// Enkripsi password baru
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to encrypt password"})
	}

	if err := db.DB.Model(&models.User{}).Where("id = ?", userId).Update("password", encryptedPassword).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password updated successfully"})
}

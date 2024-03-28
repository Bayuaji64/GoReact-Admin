package controllers

import (
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/middlewares"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Alluser(c *fiber.Ctx) error {

	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err

	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	// limit := 5
	// offset := (page - 1) * limit
	// var total int64
	// var users []models.User

	// db.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	// db.DB.Model(&models.User{}).Count(&total)

	// return c.JSON(users)

	return c.JSON(models.Paginate(db.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {

	// "email":"cc@mail.com",
	// "password":"3"

	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err

	}
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

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
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err

	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.DB.Preload("Role").Find(&user)

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

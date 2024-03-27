package controllers

import (
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Allproduct(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	// return c.JSON(users)

	return c.JSON(models.Paginate(db.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := db.DB.Create(&product).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create product"})
	}

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.DB.Find(&product)

	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	db.DB.Model(&product).Updates(product)

	return c.JSON(product)

}

func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.DB.Delete(&product)
	return c.JSON(fiber.Map{"message": "Delete user success"})

}

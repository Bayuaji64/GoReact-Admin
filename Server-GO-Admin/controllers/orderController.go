package controllers

import (
	"encoding/csv"
	"os"
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Allorder(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	// return c.JSON(users)

	return c.JSON(models.Paginate(db.DB, &models.Order{}, page))
}

func CreateOrder(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := db.DB.Create(&product).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create product"})
	}

	return c.JSON(product)
}

func GetOrder(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.DB.Find(&product)

	return c.JSON(product)

}

func UpdateOrder(c *fiber.Ctx) error {

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

func DeleteOrder(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.DB.Delete(&product)
	return c.JSON(fiber.Map{"message": "Delete user success"})

}

func Export(c *fiber.Ctx) error {
	filePath := "./csv/orders.csv"

	if err := CreateFile(filePath); err != nil {
		return err
	}

	return c.Download(filePath)

}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order

	db.DB.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID", "Name", "Product Title", "Price", "Quantity",
	})

	for _, v := range orders {

		data := []string{
			strconv.Itoa(int(v.Id)),
			v.Firstanme + " " + v.Lastname,
			v.Email,
			"",
			"",
			"",
		}

		if err := writer.Write(data); err != nil {
			return err
		}

		for _, oi := range v.OrderItems {

			data := []string{
				"",
				"",
				"",
				oi.ProductTitle,
				strconv.Itoa(int(oi.Price)),
				strconv.Itoa(int(oi.Quantity)),
			}

			if err := writer.Write(data); err != nil {
				return err
			}

		}
	}
	return nil

}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *fiber.Ctx) error {

	var sales []Sales

	db.DB.Raw(`
	SELECT DATE_FORMAT(o.created_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
	FROM orders o
	JOIN order_itmes oi on o.id = oi.order_id
	GROUP BY date
	`).Scan(&sales)

	return c.JSON(sales)

}

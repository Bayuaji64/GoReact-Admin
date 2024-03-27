package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {

	limit := 15
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)

	total := entity.Count(db)

	// return c.JSON(users)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	}
}

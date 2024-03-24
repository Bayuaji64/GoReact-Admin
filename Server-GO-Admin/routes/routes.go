package routes

import (
	"example.com/go-admin/controllers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	app.Get("/other", controllers.Other)

}

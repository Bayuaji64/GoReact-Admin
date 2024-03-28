package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	orderGroup := app.Group("/api")
	orderGroup.Use(middlewares.IsAuthenticated)

	orderGroup.Get("/orders", controllers.Allorder)
	orderGroup.Post("/export", controllers.Export)
	orderGroup.Get("/chart", controllers.Chart)

}

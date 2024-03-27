package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(app *fiber.App) {
	roleGroup := app.Group("/api")
	roleGroup.Use(middlewares.IsAuthenticated)

	roleGroup.Post("/upload", controllers.Upload)
	roleGroup.Static("/uploads", "./uploads")

}

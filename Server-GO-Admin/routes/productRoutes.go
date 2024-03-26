package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	productGroup := app.Group("/api")
	productGroup.Use(middlewares.IsAuthenticated)

	productGroup.Get("/product", controllers.Allproduct)
	productGroup.Post("/product", controllers.CreateProduct)
	productGroup.Get("/product/:id", controllers.GetProduct)
	productGroup.Put("/product/:id", controllers.UpdateProduct)
	productGroup.Delete("/product/:id", controllers.DeleteProduct)
}

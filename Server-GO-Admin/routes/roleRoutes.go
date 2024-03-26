package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	roleGroup := app.Group("/api")
	roleGroup.Use(middlewares.IsAuthenticated)

	roleGroup.Get("/roles", controllers.Allrole)
	roleGroup.Post("/roles", controllers.CreateRole)
	roleGroup.Get("/roles/:id", controllers.GetRole)
	roleGroup.Put("/roles/:id", controllers.UpdateRole)
	roleGroup.Delete("/roles/:id", controllers.DeleteRole)
}

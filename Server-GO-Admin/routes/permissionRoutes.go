package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PermissionRoutes(app *fiber.App) {
	roleGroup := app.Group("/api")
	roleGroup.Use(middlewares.IsAuthenticated)

	roleGroup.Get("/permissions", controllers.Allpermission)

}

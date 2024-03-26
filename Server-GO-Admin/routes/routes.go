package routes

import (
	"example.com/go-admin/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	UserRoutes(app)
	RoleRoutes(app)
	PermissionRoutes(app)

	app.Post("/api/logout", controllers.Logout)

}

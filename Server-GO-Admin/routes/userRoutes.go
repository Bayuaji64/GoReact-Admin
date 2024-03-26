package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/api")
	userGroup.Use(middlewares.IsAuthenticated)

	userGroup.Get("/user", controllers.User)
	userGroup.Get("/users", controllers.Alluser)
	userGroup.Post("/users", controllers.CreateUser)
	userGroup.Put("/users/info", controllers.UpdateInfo)
	userGroup.Put("/users/password", controllers.UpdatePassword)
	userGroup.Get("/users/:id", controllers.GetUser)
	userGroup.Put("/users/:id", controllers.UpdateUser)
	userGroup.Delete("/users/:id", controllers.DeleteUser)
}

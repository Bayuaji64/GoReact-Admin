package routes

import (
	"example.com/go-admin/controllers"
	"example.com/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	authRoutes := app.Group("/api")
	authRoutes.Use(middlewares.IsAuthenticated)

	authRoutes.Get("/user", controllers.User)
	authRoutes.Post("/logout", controllers.Logout)

	authRoutes.Get("/users", controllers.Alluser)
	authRoutes.Post("/users", controllers.CreateUser)
	authRoutes.Get("/users/:id", controllers.GetUser)
	authRoutes.Put("/users/:id", controllers.UpdateUser)
	authRoutes.Delete("/users/:id", controllers.DeleteUser)

}

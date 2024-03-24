package main

import (
	"log"
	"os"

	"example.com/go-admin/db"
	"example.com/go-admin/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

}

func main() {

	db.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
	}))

	routes.Setup(app)

	port := os.Getenv("PORT")

	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

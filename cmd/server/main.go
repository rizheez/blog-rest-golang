package main

import (
	"blog-rest/internal/database"
	"blog-rest/internal/models"
	httpTransport "blog-rest/internal/transport/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	database.DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Category{})
	httpTransport.SetupRoutes(app)
	app.Listen(":8000")
}

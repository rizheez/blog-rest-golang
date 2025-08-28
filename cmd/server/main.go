package main

import (
	"blog-rest/internal/database"
	"blog-rest/internal/models"
	httpTransport "blog-rest/internal/transport/http"
	"blog-rest/internal/validation"

	"github.com/gofiber/fiber/v2"
)

func main() {
	validation.InitValidator()
	app := fiber.New()

	database.Connect()

	database.DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Category{})
	httpTransport.SetupRoutes(app, database.DB)
	app.Listen(":8000")
}

package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/posts", GetPosts)
	app.Post("/posts", CreatePost)
}

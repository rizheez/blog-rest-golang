package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/posts", GetPosts)
	app.Post("/posts", CreatePost)
	app.Get("/users", GetAllUsers)
	app.Post("/users", CreateUser)
	app.Get("/categories", GetAllCategories)
	app.Post("/categories", CreateCategory)
	app.Delete("/categories/:id", DeleteCategory)
	app.Get("/categories/:id", GetCategoriesByid)
	app.Put("/categories/:id", UpdateCategory)
}

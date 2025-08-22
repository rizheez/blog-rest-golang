package http

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// app.Get("/posts", GetPosts)
	// app.Post("/posts", CreatePost)
	// app.Get("/users", GetAllUsers)
	// app.Post("/users", CreateUser)
	// app.Get("/categories", GetAllCategories)
	// app.Post("/categories", CreateCategory)
	// app.Delete("/categories/:id", DeleteCategory)
	// app.Get("/categories/:id", GetCategoriesByid)
	// app.Put("/categories/:id", UpdateCategory)

	// user
	user := app.Group("/users")
	user.Get("/", GetAllUsers)
	user.Post("/", CreateUser)
	user.Get("/:id", GetUserById)
	user.Put("/:id", UpdateUser)
	user.Delete("/:id", DeleteUser)

	// category
	category := app.Group("/categories")
	category.Get("/", GetAllCategories)
	category.Post("/", CreateCategory)
	category.Get("/:id", GetCategoriesByid)
	category.Delete("/:id", DeleteCategory)
	category.Put("/:id", UpdateCategory)

	// post
	post := app.Group("/post")
	post.Get("/", GetPosts)
	post.Post("/", CreatePost)
}

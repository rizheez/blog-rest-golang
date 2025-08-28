package http

import (
	"blog-rest/internal/middleware"
	"blog-rest/internal/repository"
	"blog-rest/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// app.Get("/posts", GetPosts)
	// app.Post("/posts", CreatePost)
	// app.Get("/users", GetAllUsers)
	// app.Post("/users", CreateUser)
	// app.Get("/categories", GetAllCategories)
	// app.Post("/categories", CreateCategory)
	// app.Delete("/categories/:id", DeleteCategory)
	// app.Get("/categories/:id", GetCategoriesByid)
	// app.Put("/categories/:id", UpdateCategory)
	// login and register
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)
	authService := services.NewAuthService(userRepo)
	authHandler := NewAuthHandler(authService)
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	// user
	user := app.Group("/users")
	user.Use(middleware.JWTProtected(userRepo))
	user.Get("/me", userHandler.GetMe)
	user.Get("/", userHandler.GetAllUsers)
	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUserById)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)

	// category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)
	category := app.Group("/categories")
	category.Use(middleware.JWTProtected(userRepo))
	category.Get("/", categoryHandler.GetAllCategories)
	category.Post("/", categoryHandler.CreateCategory)
	category.Get("/:id", categoryHandler.GetCategoriesByid)
	category.Delete("/:id", categoryHandler.DeleteCategory)
	category.Put("/:id", categoryHandler.UpdateCategory)

	// post
	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := NewPostHandler(postService)
	post := app.Group("/posts")
	post.Use(middleware.JWTProtected(userRepo))
	post.Get("/", postHandler.GetPosts)
	post.Post("/", postHandler.CreatePost)
	post.Get("/:id", postHandler.GetPostById)
	post.Put("/:id", postHandler.UpdatePost)
	post.Delete("/:id", postHandler.DeletePost)
}

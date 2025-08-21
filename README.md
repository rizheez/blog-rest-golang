# Blog REST API

Blog REST API is a simple blog application built with **Go**, using the **Fiber framework** for HTTP handling and **GORM** as the ORM.

## 🚀 Features

- User management (list users, create user)
- Category management (list categories, create category)
- Post management (list posts, create post)
- RESTful API design with clear separation of concerns (Repository → Service → Handler)

## 📂 Project Structure

```
.
├── cmd/
│   └── server/            # Application entry point (main.go)
├── internal/
│   ├── repository/        # Database layer
│   │   ├── category_repository.go
│   │   └── user_repository.go
│   ├── services/          # Business logic layer
│   │   ├── category_service.go
│   │   └── user_service.go
│   └── transport/
│       └── http/          # HTTP handlers & routes
│           ├── category_handler.go
│           ├── user_handler.go
│           └── routes.go
└── go.mod
```

### Layers

- **Repository** → Database interaction (CRUD queries)
- **Service** → Business logic between repository and handlers
- **Handler (HTTP)** → HTTP request handling & response mapping

## 📌 API Endpoints

### Categories

- `GET /categories` → Fetch all categories
- `POST /categories` → Create a new category

### Users

- `GET /users` → Fetch all users
- `POST /users` → Create a new user

### Posts

- `GET /posts` → Fetch all posts
- `POST /posts` → Create a new post

## ⚙️ Tech Stack

- [Go](https://go.dev/) — Programming language
- [Fiber](https://gofiber.io/) — Web framework
- [GORM](https://gorm.io/) — ORM for database handling

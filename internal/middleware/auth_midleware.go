package middleware

import (
	"strings"

	"blog-rest/internal/repository"
	"blog-rest/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// Optionally inject userRepo if you want to fetch user object in middleware
func JWTProtected(userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1) get token from Authorization header "Bearer <token>"
		authHeader := c.Get("Authorization")
		var token string
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// fallback: get from cookie
			token = c.Cookies("jwt")
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "message": "Please log in to access this feature", "code": fiber.StatusUnauthorized})
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "message": "Your session has expired or is invalid. Please log in again", "code": fiber.StatusUnauthorized})
		}

		// optionally fetch user and set to locals
		if userRepo != nil {
			user, err := userRepo.GetUserById(claims.UserID)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "message": "Your account could not be found. Please contact support or try logging in again", "code": fiber.StatusUnauthorized})
			}
			user.Password = "" // avoid leaking
			c.Locals("user", user)
		} else {
			c.Locals("user_id", claims.UserID)
		}

		return c.Next()
	}
}

// Simple JWT middleware without user repository dependency
func JWTRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1) get token from Authorization header "Bearer <token>"
		authHeader := c.Get("Authorization")
		var token string
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// fallback: get from cookie
			token = c.Cookies("jwt")
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Please log in to access this feature", "code": fiber.StatusUnauthorized})
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Your session has expired or is invalid. Please log in again", "code": fiber.StatusUnauthorized})
		}

		// Set user ID in context for handlers to use
		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}

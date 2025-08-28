package http

import (
	"blog-rest/internal/dto"
	"blog-rest/internal/services"
	"blog-rest/internal/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
}
type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func NewAuthHandler(a services.AuthService) AuthHandler {
	return &authHandler{authService: a}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var input dto.RegisterDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
			"code":  fiber.StatusBadRequest,
		})
	}

	// Validate input
	if err := validation.Validate.Struct(&input); err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg := make(map[string]string)
		for _, e := range errs {
			field := e.Field()
			errMsg[field] = e.Translate(validation.Trans)
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errMsg,
			"code":  fiber.StatusBadRequest,
		})
	}

	user, err := h.authService.Register(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToUserResponse(user))
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var input dto.LoginDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
			"code":  fiber.StatusBadRequest,
		})
	}

	// Validate input
	if err := validation.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed: " + err.Error(),
			"code":  fiber.StatusBadRequest,
		})
	}

	token, err := h.authService.Login(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.TokenDTO{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
	})

}

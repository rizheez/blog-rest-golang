package http

import (
	"blog-rest/internal/dto"
	"blog-rest/internal/models"
	"blog-rest/internal/services"
	"blog-rest/internal/validation"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	GetAllUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	GetMe(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

func NewUserHandler(u services.UserService) UserHandler {
	return &userHandler{userService: u}
}

type userHandler struct {
	userService services.UserService
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	UserResponse := make([]dto.UserResponse, len(users))
	for i, u := range users {
		UserResponse[i] = dto.ToUserResponse(&u)
	}
	return c.Status(fiber.StatusOK).JSON(UserResponse)

}

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToUserResponse(user))
}

// GetMe gets the current user's profile based on JWT token
func (h *userHandler) GetMe(c *fiber.Ctx) error {
	// Get user from JWT middleware context
	user := c.Locals("user")
	if user != nil {
		// If full user object is available from JWTProtected middleware
		userObj := user.(*models.User)
		return c.Status(fiber.StatusOK).JSON(dto.ToUserResponse(userObj))
	}

	// If only user_id is available from JWTRequired middleware
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unable to identify user from token",
			"code":  fiber.StatusUnauthorized,
		})
	}

	// Fetch user by ID
	id, ok := userID.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID format",
			"code":  fiber.StatusInternalServerError,
		})
	}

	userObj, err := h.userService.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Your profile could not be found",
			"code":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToUserResponse(userObj))
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	if err := validation.Validate.Struct(req); err != nil {
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	if err := h.userService.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToUserResponse(&user))
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	user.ID = uint(id)
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
				"code":  fiber.StatusInternalServerError,
			})
		}
		user.Password = string(hashedPassword)
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	if err := h.userService.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToUserResponse(&user))

}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var user models.User
	user.ID = uint(id)
	err = h.userService.DeleteUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

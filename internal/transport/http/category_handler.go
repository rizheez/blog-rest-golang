package http

import (
	"blog-rest/internal/dto"
	"blog-rest/internal/models"
	"blog-rest/internal/services"
	"blog-rest/internal/validation"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler interface {
	GetAllCategories(c *fiber.Ctx) error
	GetCategoriesByid(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}
type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(c services.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: c}
}

func (h *categoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	CategoryResponse := make([]dto.CategoryResponse, len(categories))
	for i, c := range categories {
		CategoryResponse[i] = dto.ToCategoryResponse(&c)
	}
	return c.Status(fiber.StatusOK).JSON(CategoryResponse)
}

func (h *categoryHandler) GetCategoriesByid(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	category, err := h.categoryService.GetCategoriesByid(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToCategoryResponse(category))
}

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
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

	category := models.Category{
		Name: req.Name,
	}
	if err := h.categoryService.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToCategoryResponse(&category))
}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}

	var req dto.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	category := models.Category{
		Name: req.Name,
	}
	category.ID = uint(id)
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	if err := h.categoryService.UpdateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToCategoryResponse(&category))
}

func (h *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var category models.Category
	category.ID = uint(id)
	err = h.categoryService.DeleteCategory(&category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

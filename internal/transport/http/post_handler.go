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

type PostHandler interface {
	GetPosts(c *fiber.Ctx) error
	GetPostById(c *fiber.Ctx) error
	CreatePost(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}

type postHandler struct {
	postService services.PostService
}

func NewPostHandler(ps services.PostService) PostHandler {
	return &postHandler{postService: ps}
}

func (h *postHandler) GetPosts(c *fiber.Ctx) error {
	posts, err := h.postService.GetPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	PostResponse := make([]dto.PostResponse, len(posts))
	for i, p := range posts {
		PostResponse[i] = dto.ToPostResponse(p)
	}
	return c.Status(fiber.StatusOK).JSON(PostResponse)
}

func (h *postHandler) GetPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	post, err := h.postService.GetPostById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToPostResponse(post))
}

func (h *postHandler) CreatePost(c *fiber.Ctx) error {
	var req dto.CreatePostRequest
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
	post := models.Post{
		Title:      req.Title,
		Body:       req.Body,
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
	}

	if err := h.postService.CreatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToPostResponse(post))

}

func (h *postHandler) UpdatePost(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var req dto.UpdatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	post := models.Post{
		Title:      req.Title,
		Body:       req.Body,
		CategoryID: req.CategoryID,
	}
	post.ID = uint(id)
	if err := h.postService.UpdatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToPostResponse(post))
}

func (h *postHandler) DeletePost(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var post models.Post
	post.ID = uint(id)
	err = h.postService.DeletePost(&post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

package http

import (
	"blog-rest/internal/models"
	"blog-rest/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	posts, err := services.GetPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(posts)
}

func GetPostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	post, err := services.GetPostById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusBadRequest,
		})
	}
	if err := services.CreatePost(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}

	return c.JSON(post)

}

func UpdatePost(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id",
			"code":  fiber.StatusBadRequest,
		})
	}
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	post.ID = uint(id)
	if err := services.UpdatePost(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
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
	err = services.DeletePost(&post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"code":  fiber.StatusInternalServerError,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

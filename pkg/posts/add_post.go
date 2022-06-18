package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

type AddPostRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
}

func (h handler) AddPost(c *fiber.Ctx) error {
	body := AddPostRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var post models.Post

	post.Title = body.Title
	post.Description = body.Description
	post.CategoryId = body.CategoryId

	if result := h.DB.Create(&post); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&post)
}

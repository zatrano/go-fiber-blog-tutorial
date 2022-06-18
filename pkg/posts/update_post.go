package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

type UpdatePostRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
}

func (h handler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdatePostRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var post models.Post

	if result := h.DB.First(&post, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	post.Title = body.Title
	post.Description = body.Description
	post.CategoryId = body.CategoryId

	// save post
	h.DB.Save(&post)

	return c.Status(fiber.StatusOK).JSON(&post)
}

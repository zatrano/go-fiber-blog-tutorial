package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

func (h handler) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post

	if result := h.DB.First(&post, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&post)
}

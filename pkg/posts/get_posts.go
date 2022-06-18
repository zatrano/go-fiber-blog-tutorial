package posts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

func (h handler) GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	if result := h.DB.Find(&posts); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&posts)
}

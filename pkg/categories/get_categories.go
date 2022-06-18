package categories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

func (h handler) GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	if result := h.DB.Find(&categories); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
}

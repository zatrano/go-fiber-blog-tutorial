package categories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

func (h handler) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&category)
}

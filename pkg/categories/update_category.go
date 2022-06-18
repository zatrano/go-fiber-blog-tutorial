package categories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

type UpdateCategoryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateCategoryRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	category.Title = body.Title
	category.Description = body.Description

	// save category
	h.DB.Save(&category)

	return c.Status(fiber.StatusOK).JSON(&category)
}

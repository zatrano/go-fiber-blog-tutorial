package categories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/models"
)

type AddCategoryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h handler) AddCategory(c *fiber.Ctx) error {
	body := AddCategoryRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	category.Title = body.Title
	category.Description = body.Description

	if result := h.DB.Create(&category); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&category)
}

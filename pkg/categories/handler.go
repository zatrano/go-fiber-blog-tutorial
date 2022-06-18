package categories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/categories")
	routes.Post("/", h.AddCategory)
	routes.Get("/", h.GetCategories)
	routes.Get("/:id", h.GetCategory)
	routes.Put("/:id", h.UpdateCategory)
	routes.Delete("/:id", h.DeleteCategory)
}

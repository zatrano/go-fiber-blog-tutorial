package posts

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

	routes := app.Group("/posts")
	routes.Post("/", h.AddPost)
	routes.Get("/", h.GetPosts)
	routes.Get("/:id", h.GetPost)
	routes.Put("/:id", h.UpdatePost)
	routes.Delete("/:id", h.DeletePost)
}

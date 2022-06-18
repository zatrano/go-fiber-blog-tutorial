package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zatrano/go-blog-tutorial/common/config"
	"github.com/zatrano/go-blog-tutorial/common/db"
	"github.com/zatrano/go-blog-tutorial/pkg/categories"
	"github.com/zatrano/go-blog-tutorial/pkg/posts"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	posts.RegisterRoutes(app, db)
	categories.RegisterRoutes(app, db)

	app.Listen(c.Port)
}

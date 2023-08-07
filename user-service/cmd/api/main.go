package main

import (
	"github.com/brandon-a-pinto/nebula/user-service/internal/db"
	"github.com/brandon-a-pinto/nebula/user-service/internal/handler"
	"github.com/gofiber/fiber/v2"
)

var (
	h = handler.NewUserHandler()
)

func main() {
	db.PostgreSQLConnection()
	app := fiber.New()

	v1 := app.Group("/api/v1")
	{
		v1.Post("/", h.HandlePostUser)
	}

	app.Listen(":80")
}

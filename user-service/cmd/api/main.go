package main

import (
	"github.com/brandon-a-pinto/nebula/user-service/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.PostgreSQLConnection()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("[GET] User Service")
	})

	app.Listen(":80")
}

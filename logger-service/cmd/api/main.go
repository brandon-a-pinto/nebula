package main

import (
	"github.com/brandon-a-pinto/nebula/logger-service/internal/db"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/handler"
	"github.com/gofiber/fiber/v2"
)

var (
	h = handler.NewLoggerHandler()
)

func main() {
	db.MongoDBConnection()
	app := fiber.New()

	app.Post("/", h.HandleLogger)

	app.Listen(":80")
}

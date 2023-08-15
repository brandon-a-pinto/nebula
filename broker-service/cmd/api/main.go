package main

import (
	"github.com/brandon-a-pinto/nebula/broker-service/internal/handler"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

var (
	h = handler.NewBrokerHandler()
)

func main() {
	rabbitmq.RabbitMQConnection()
	app := fiber.New()

	app.Post("/", h.HandleBroker)

	app.Listen(":80")
}

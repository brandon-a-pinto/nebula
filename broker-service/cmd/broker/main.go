package main

import (
	"fmt"

	"github.com/brandon-a-pinto/nebula/broker-service/configs"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/mq"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/web"
)

func main() {
	config := configs.LoadConfig()

	// RabbitMQ
	rabbitmq := mq.NewRabbitMQServer(config.ListenerRabbitmq)
	fmt.Println("Establishing RabbitMQ connection...")
	go rabbitmq.Start()

	// Web Server
	webserver := web.NewWebServer(":" + config.WebServerPort)
	fmt.Println("Starting web server on port", config.WebServerPort)
	webserver.Start()
}

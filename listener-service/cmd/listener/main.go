package main

import (
	"log"

	"github.com/brandon-a-pinto/nebula/listener-service/internal/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@listener-rabbitmq")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := event.NewConsumer(conn)
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARN", "log.ERROR"})
	if err != nil {
		log.Fatal(err)
	}
}

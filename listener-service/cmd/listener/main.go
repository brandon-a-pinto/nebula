package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@listener-rabbitmq")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	defer conn.Close()
}

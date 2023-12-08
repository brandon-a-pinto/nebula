package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQServer struct {
	Address string
}

var RMQ *amqp.Channel

func NewRabbitMQServer(addr string) *RabbitMQServer {
	return &RabbitMQServer{
		Address: addr,
	}
}

func (s *RabbitMQServer) setup(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		"logs",  // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		queue.Name, // name
		"log.*",    // key
		"logs",     // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *RabbitMQServer) Start() {
	conn, err := amqp.Dial(s.Address)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	s.setup(ch)

	RMQ = ch
}

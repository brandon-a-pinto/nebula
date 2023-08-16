package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type Request struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

func (c *Consumer) Listen(topics []string) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	loop := make(chan bool)
	go func() {
		for d := range msgs {
			var req Request
			_ = json.Unmarshal(d.Body, &req)
			go handleRequest(req)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", q.Name)
	<-loop
	return nil
}

func handleRequest(req Request) {
	switch req.Name {
	case "log", "event":
		err := logEvent(req)
		if err != nil {
			log.Println(err)
		}
	default:
		err := logEvent(req)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(data Request) error {
	jsonValue, _ := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, "http://logger-service", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (c *Consumer) setup() error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(ch)
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	if err := consumer.setup(); err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

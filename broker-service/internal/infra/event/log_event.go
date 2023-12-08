package event

import (
	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type LogEvent struct{}

func NewLogEvent() *LogEvent {
	return &LogEvent{}
}

func (e *LogEvent) Create(event, key string) error {
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(event),
	}

	err := mq.RMQ.Publish(
		"logs",
		key,
		false,
		false,
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}

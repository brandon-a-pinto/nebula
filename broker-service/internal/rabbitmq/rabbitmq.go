package rabbitmq

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RMQ RabbitMQInstance

type RabbitMQInstance struct {
	Conn *amqp.Connection
}

func RabbitMQConnection() {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		conn, err := amqp.Dial(os.Getenv("LISTENER_RABBITMQ"))
		if err != nil {
			fmt.Println("RabbitMQ not ready yet...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = conn
			break
		}

		if counts > 8 {
			log.Fatal(err)
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}

	RMQ = RabbitMQInstance{
		Conn: connection,
	}
}

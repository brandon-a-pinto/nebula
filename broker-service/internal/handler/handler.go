package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/event"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/models"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

type Broker struct{}

func NewBrokerHandler() *Broker {
	return &Broker{}
}

func (h *Broker) HandleBroker(c *fiber.Ctx) error {
	params := new(models.Request)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg":     "BodyParser error",
			"error":   err,
			"success": false,
		})
	}

	switch params.Action {
	case "user_create":
		res, err := createUserRequest(params.Data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":     "Request failed",
				"error":   err,
				"success": false,
			})
		}
		return c.JSON(res)
	case "log":
		err := logWithRabbitMQ(params.Log)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":     "Request failed",
				"error":   err,
				"success": false,
			})
		}
		return c.JSON("Logged successfully")
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "Unknown action",
			"error":   "server could not identify action",
			"success": false,
		})
	}
}

func createUserRequest(params models.CreateUserParams) (interface{}, error) {
	jsonValue, _ := json.Marshal(params)
	req, err := http.NewRequest(http.MethodPost, "http://user-service/api/v1", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var jsonData interface{}
	json.NewDecoder(res.Body).Decode(&jsonData)
	return jsonData, nil
}

func logItem(params models.LogParams) (interface{}, error) {
	jsonValue, _ := json.Marshal(params)
	req, err := http.NewRequest(http.MethodPost, "http://logger-service", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var jsonData interface{}
	json.NewDecoder(res.Body).Decode(&jsonData)
	return jsonData, nil
}

func logWithRabbitMQ(params models.LogParams) error {
	err := pushToQueue(params.Name, params.Data)
	if err != nil {
		return err
	}

	return nil
}

func pushToQueue(name string, data any) error {
	emitter, err := event.NewEventEmitter(rabbitmq.RMQ.Conn)
	if err != nil {
		return err
	}

	req := models.LogParams{
		Name: name,
		Data: data,
	}

	jsonValue, _ := json.Marshal(req)
	err = emitter.Push(string(jsonValue), "log.INFO")
	if err != nil {
		return err
	}

	return nil
}

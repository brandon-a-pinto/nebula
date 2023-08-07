package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/models"
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
	case "create_user":
		res, err := createUserRequest(params.Data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":     "Request failed",
				"error":   err,
				"success": false,
			})
		}
		return c.JSON(res)
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

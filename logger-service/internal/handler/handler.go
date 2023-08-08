package handler

import (
	"time"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/db"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Logger struct{}

func NewLoggerHandler() *Logger {
	return &Logger{}
}

func (h *Logger) HandleLogger(c *fiber.Ctx) error {
	params := new(models.Request)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":     "BodyParser error",
			"error":   err,
			"success": false,
		})
	}

	collection := db.MI.Collection(params.Name)
	log := models.Log{
		Data:      params.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := collection.InsertOne(c.Context(), log)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":     "Database error",
			"error":   err,
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"msg":     "Logged successfully",
		"data":    res,
		"success": true,
	})
}

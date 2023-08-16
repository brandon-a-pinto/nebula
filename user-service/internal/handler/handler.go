package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brandon-a-pinto/nebula/user-service/internal/db"
	"github.com/brandon-a-pinto/nebula/user-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func NewUserHandler() *User {
	return &User{}
}

func (h *User) HandlePostUser(c *fiber.Ctx) error {
	params := new(models.CreateUserParams)

	// Parse params with BodyParser
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg":     "BodyParser error",
			"error":   err,
			"success": false,
		})
	}

	// Generate encrypted password
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":     "Bcrypt error",
			"error":   err,
			"success": false,
		})
	}

	// Postgres query
	user := models.User{
		Email:       params.Email,
		Username:    params.Username,
		DisplayName: params.DisplayName,
		Password:    string(encpw),
		IsAdmin:     false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	query := `
		INSERT INTO users
		(email, username, display_name, pwd, is_admin, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	res, err := db.PI.Postgres.Query(
		query,
		user.Email,
		user.Username,
		user.DisplayName,
		user.Password,
		user.IsAdmin,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":     "Database error",
			"error":   err,
			"success": false,
		})
	}

	// Log operation
	err = logRequest("log", models.LogParams{Name: "users", Data: fmt.Sprintf("Account created: %s", params.Email)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":     "Logger error",
			"error":   err,
			"success": false,
		})
	}

	// Success response
	return c.JSON(fiber.Map{
		"msg":     "User created successfully",
		"data":    res,
		"success": true,
	})
}

func logRequest(action string, log models.LogParams) error {
	var request struct {
		Action string           `json:"action"`
		Log    models.LogParams `json:"log"`
	}

	request.Action = action
	request.Log = log

	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, "http://broker-service", bytes.NewBuffer(jsonValue))
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

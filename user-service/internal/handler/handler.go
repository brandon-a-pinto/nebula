package handler

import (
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

	// Generates encrypted password
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

	// Success response
	return c.JSON(fiber.Map{
		"msg":     "User created successfully",
		"data":    res,
		"success": true,
	})
}

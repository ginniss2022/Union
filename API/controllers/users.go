package controllers

import (
	initializer "github.com/ginniss2022/union/config"
	"github.com/ginniss2022/union/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type createUserRequest struct {
	Email    string
	Password string
}

func CreateNewUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user := models.User{Email: req.Email, Password: string(hash)}
	result := initializer.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.SendStatus(200)
}

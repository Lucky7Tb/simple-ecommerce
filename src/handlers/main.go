package handlers

import (
	"simple-ecommerce/src/config"
	"simple-ecommerce/src/handlers/auth"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewHandler(route *fiber.App, db *gorm.DB) {
	var validate = config.Validator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	auth.NewAuthHandler(route, db, &validate)
}

package main

import (
	"fmt"
	"os"
	"simple-ecommerce/src/config"
	"simple-ecommerce/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := config.ConnectToDatabase()
	app := fiber.New()
	handlers.NewHandler(app, db)
	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))

}

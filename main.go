package main

import (
	"fiber-api/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
)

func main()  {

	// load env
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error locading .env: %w \n", err))
	}

	config.Connect()

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"info": "GoLang, Fiber, Postgres API",
		})
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Errorf("Failed to listen on port %s: %w \n", os.Getenv("PORT"), err))
	}
}
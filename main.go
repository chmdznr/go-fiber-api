package main

import (
	"fiber-api/config"
	"fiber-api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error locading .env: %w \n", err))
	}

	// connect DB
	config.DbConnect()

	// create app
	app := fiber.New()

	// sample middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"info": "GoLang, Fiber, Postgres API",
		})
	})

	// config router
	routes.Setup(app)

	// listen
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Errorf("Failed to listen on port %s: %w \n", os.Getenv("PORT"), err))
	}
}

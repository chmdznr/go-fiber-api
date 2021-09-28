package routes

import (
	"fiber-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUsersById)
	app.Post("/users", controllers.CreateUsers)
	app.Put("/users/:id", controllers.UpdateUsers)
	app.Delete("/users/:id", controllers.DeleteUsers)
}

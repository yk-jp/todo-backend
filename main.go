package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yk-jp/todo-backend/config"
	"github.com/yk-jp/todo-backend/database"
	"github.com/yk-jp/todo-backend/routes"
)

func setupRoute(app *fiber.App) {
	// task
	app.Get("/api/task", routes.GetTasks)
	app.Get("/api/task/:id", routes.GetTask)
	app.Post("/api/task", routes.CreateTask)
}

func main() {
	app := fiber.New()
	config := config.LoadEnvVariables()
	database.ConnectDb(config)
	setupRoute(app)
	log.Fatal(app.Listen(":5000"))
}

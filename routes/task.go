package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yk-jp/todo-backend/database"
	"github.com/yk-jp/todo-backend/database/schema"
	"github.com/yk-jp/todo-backend/models"
)

func CreateTask(c *fiber.Ctx) error {
	var task schema.Task

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Db.Db.Create(&task)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	var responseData models.Task
	response := database.Db.Db.Model(&models.Task{}).Select("tasks.id, title, status_refer, statuses.name as status").Joins("left join statuses on statuses.id = tasks.status_refer").Where("tasks.id = ?", task.ID).Scan(&responseData)

	if response.Error != nil {
		return c.Status(500).JSON("Network Error")
	}

	return c.Status(200).JSON(responseData)
}

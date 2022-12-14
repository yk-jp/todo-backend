package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yk-jp/todo-backend/database"
	"github.com/yk-jp/todo-backend/database/schema"
	"github.com/yk-jp/todo-backend/models"
	"github.com/yk-jp/todo-backend/utils"
)

func GetTasks(c *fiber.Ctx) error {
	responseData := []models.Task{}

	response := utils.FindAllTasks(&responseData)

	if response.Error != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(responseData)
}

func GetTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.ErrBadRequest
	}

	var responseData models.Task
	response := utils.FindTaskById(&responseData, id)

	if response.Error != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(responseData)
}

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
	response := utils.FindTaskById(&responseData, int(task.ID))

	if response.Error != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(responseData)
}

func UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.ErrBadRequest
	}

	var updateData schema.Task

	if err := c.BodyParser(&updateData); err != nil {
		return err
	}

	response := database.Db.Db.Model(&schema.Task{}).Where("id = ?", id).Updates(updateData)

	if response.Error != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).SendString("Successfully updated")
}

func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.ErrBadRequest
	}

	response := database.Db.Db.Delete(&schema.Task{}, id)

	if response.Error != nil {
		fmt.Println(response.Error)
		return fiber.ErrInternalServerError
	}

	return c.Status(200).SendString("Successfully deleted")
}

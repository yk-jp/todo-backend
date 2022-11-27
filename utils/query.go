package utils

import (
	"github.com/yk-jp/todo-backend/database"
	"github.com/yk-jp/todo-backend/models"
	"gorm.io/gorm"
)

func FindAllTasks(responseData *[]models.Task) *gorm.DB {
	return database.Db.Db.Model(&models.Task{}).
		Select("tasks.id, title, status_refer, statuses.name as status").
		Joins("left join statuses on statuses.id = tasks.status_refer").
		Scan(&responseData)
}

func FindTaskById(responseData *models.Task, id int) *gorm.DB {
	return database.Db.Db.Model(&models.Task{}).
		Select("tasks.id, title, status_refer, statuses.name as status").
		Joins("left join statuses on statuses.id = tasks.status_refer").
		Where("tasks.id = ?", id).
		Scan(&responseData)
}

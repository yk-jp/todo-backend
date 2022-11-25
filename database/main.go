package database

import (
	"fmt"
	"log"

	"github.com/yk-jp/todo-backend/config"
	"github.com/yk-jp/todo-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Db DBInstance

func ConnectDb(config config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=todo", config.Db.Host, config.Db.User, config.Db.Password, config.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Conecction was successful")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	db.AutoMigrate(&models.Task{}, &models.Status{})
	Db = DBInstance{Db: db}
}

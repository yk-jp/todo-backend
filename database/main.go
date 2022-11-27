package database

import (
	"fmt"
	"log"

	"github.com/yk-jp/todo-backend/config"
	"github.com/yk-jp/todo-backend/database/schema"
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

	// create enum type in postgres
	createStatus(db)

	log.Println("Running migrations")
	db.AutoMigrate(&schema.Task{}, &schema.Status{})
	Db = DBInstance{Db: db}
}

func createStatus(db *gorm.DB) {
	statuses := []schema.Status{{Name: schema.Pending}, {Name: schema.Complete}}
	db.Create(&statuses)
}

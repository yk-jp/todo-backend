package database

import (
	"fmt"
	"log"

	"github.com/yk-jp/todo-backend/config"
	"github.com/yk-jp/todo-backend/database/models"
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
	createEnum(db)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Task{}, &models.Status{})
	Db = DBInstance{Db: db}
}

func createEnum(db *gorm.DB) {
	// https://github.com/jackc/pgx/issues/498
	// must be hard-coded unless variables are safely interpolated into the sql string with a proper library.
	db.Raw("CREATE TYPE task_status AS ENUM ('pending', 'complete');")
}

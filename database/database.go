package database

import (
	"fmt"
	"log"

	"shop/config"
	"shop/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=%s port=%s sslmode=disable",
		config.Config("DATABASE_PASSWORD"),
		config.Config("DATABASE_NAME"),
		config.Config("DATABASE_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})

	Db = db
}

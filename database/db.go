package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shoe-store/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=asa123 dbname=shoestore port=5432 sslmode=disable TimeZone=Asia/Almaty"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	database.AutoMigrate(&models.Shoe{})

	DB = database
}

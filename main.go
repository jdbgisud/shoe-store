package main

import (
	"github.com/gin-gonic/gin"
	"shoe-store/database"
	"shoe-store/routes"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	database.ConnectDatabase()

	// Маршруты
	routes.SetupRoutes(r)

	// Запуск сервера
	r.Run(":8081")
}

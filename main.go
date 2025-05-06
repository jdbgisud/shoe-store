package main

import (
	"github.com/gin-gonic/gin"
	"shoe-store/database"
	"shoe-store/routes"
)

func main() {
	// Создаем новый роутер
	r := gin.Default()

	// Подключение к базе данных
	database.ConnectDatabase()

	// Настройка маршрутов
	routes.SetupRoutes(r)

	// Запуск сервера
	err := r.Run(":8081")
	if err != nil {
		panic("Не удалось запустить сервер: " + err.Error())
	}
}

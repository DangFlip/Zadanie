package main

import (
	"log"
	"todo-app/config"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Инициализация маршрутов
	r := gin.Default()
	routes.InitRoutes(r, db)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
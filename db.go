package config

import (
	"log"
	"todo-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Подключение к SQLite
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Ошибка миграции базы данных:", err)
	}

	return db, nil
}
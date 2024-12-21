package controllers

import (
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

// Создание задачи
func (t *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if err := t.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания задачи"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Получение всех задач
func (t *TaskController) GetTasks(c *gin.Context) {
	var tasks []models.Task
	status := c.Query("status")
	query := t.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения задач"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Получение задачи по ID
func (t *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := t.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Обновление задачи
func (t *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := t.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	t.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, task)
}

// Удаление задачи
func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := t.DB.Delete(&models.Task{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления задачи"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
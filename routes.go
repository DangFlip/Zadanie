package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	taskController := controllers.TaskController{DB: db}

	// Маршруты для задач
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("", taskController.CreateTask)
		taskRoutes.GET("", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
}
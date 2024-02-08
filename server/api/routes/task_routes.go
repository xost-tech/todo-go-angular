package routes

import (
	"todo-go-angular/server/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup) {
	taskRoutes := router.Group("/tasks")

	taskRoutes.GET("/", handlers.GetTasksHandler)
	taskRoutes.GET("/user/:userId", handlers.GetTasksByUserIdHandler)
	taskRoutes.GET("/:taskId", handlers.GetTaskByIdHandler)
	taskRoutes.POST("/", handlers.CreateTaskHandler)
	taskRoutes.PUT("/:taskId", handlers.UpdateTaskHandler)
	taskRoutes.DELETE("/:taskId", handlers.DeleteTaskHandler)
}

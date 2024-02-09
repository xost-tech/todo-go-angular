package routes

import (
	"todo-go-angular/server/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.GET("/", handlers.GetCategoriesHandler)
	categoryRoutes.GET("/user/:userId", handlers.GetCategoriesByUserIdHandler)
	categoryRoutes.GET("/:categoryId", handlers.GetCategoryByIdHandler)
	categoryRoutes.POST("/", handlers.CreateCategoryHandler)
	categoryRoutes.PUT("/:categoryId", handlers.UpdateCategoryHandler)
	categoryRoutes.DELETE("/:categoryId", handlers.DeleteCategoryHandler)
}

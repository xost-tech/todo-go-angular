package main

import (
	"todo-go-angular/server/api/middleware"
	"todo-go-angular/server/api/routes"
	"todo-go-angular/server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	routes.SetupTaskRoutes(router.Group("/api"))

	router.Run(":8080")
}

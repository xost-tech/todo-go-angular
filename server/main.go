package main

import (
	"todo-go-angular/server/models"
)

func main() {
	models.InitDB()
	defer models.CloseDB()
}

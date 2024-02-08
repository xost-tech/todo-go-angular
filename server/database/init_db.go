package database

import (
	"database/sql"
	"fmt"
	"todo-go-angular/server/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dbConfig := config.LoadDBConfig()

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to ping database")
	}

	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

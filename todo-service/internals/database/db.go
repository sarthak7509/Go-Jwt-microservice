package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS todo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  
    user_id INTEGER NOT NULL,             
    title TEXT NOT NULL,                  
    description TEXT,                     
    status TEXT NOT NULL DEFAULT 'pending', 
    priority TEXT DEFAULT 'medium',       
    due_date DATETIME,                   
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP  
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Print(err)
		panic("Oops error")
	}

}

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	//_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create events table.")
	}
}

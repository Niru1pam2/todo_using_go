package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Error connecting to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	isFinished BOOLEAN,
	createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(fmt.Sprintf("Failed to create tasks table: %v", err))
	}

	createTrigger := `
	CREATE TRIGGER IF NOT EXISTS trigger_tasks_updatedAt 
	AFTER UPDATE ON tasks
	FOR EACH ROW
	BEGIN
		UPDATE tasks SET updatedAt = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;`

	_, err = DB.Exec(createTrigger)
	if err != nil {
		panic(fmt.Sprintf("Failed to create trigger: %v", err))
	}

	fmt.Println("Database tables and triggers initialized successfully!")
}

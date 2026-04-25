package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type db struct {
	db *sql.DB
}

func InitDB(dataSourceName string) (*sql.DB, error) {
	// Open the database connection
	dbConn, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)

	// sql.Open doesn't actually connect immediately. We ping it to verify.
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}

	// Run migrations
	if err := createTables(dbConn); err != nil {
		return nil, err
	}

	fmt.Println("Database connection and tables initialized successfully!")
	return dbConn, nil
}

func createTables(dbConn *sql.DB) error {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	isFinished BOOLEAN,
	createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := dbConn.Exec(createUsersTable)

	if err != nil {
		return err
	}

	createTrigger := `
	CREATE TRIGGER IF NOT EXISTS trigger_tasks_updatedAt 
	AFTER UPDATE ON tasks
	FOR EACH ROW
	BEGIN
		UPDATE tasks SET updatedAt = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;`

	_, err = dbConn.Exec(createTrigger)
	if err != nil {
		return err
	}

	fmt.Println("Database tables and triggers initialized successfully!")

	return nil
}

package models

import (
	"database/sql"
	"errors"
	"time"
	"todo_app/db"
)

type Task struct {
	Id         int64
	Title      string `binding:"required"`
	IsFinished bool
	createdAt  time.Time
	updatedAt  time.Time
}

var Tasks []Task

func GetAllTasks() ([]Task, error) {
	query := "SELECT * FROM tasks"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Tasks []Task

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.Id, &task.Title, &task.IsFinished, &task.createdAt, &task.updatedAt); err != nil {
			return nil, err
		}

		Tasks = append(Tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Tasks, nil
}

func (t *Task) GetTask() error {

	query := "SELECT * FROM tasks WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(t.Id).Scan(&t.Id, &t.Title, &t.IsFinished, &t.createdAt, &t.updatedAt)

	if err != nil {
		return err
	}

	if err != nil {
		// 3. Check if the error is specifically because the ID doesn't exist
		if err == sql.ErrNoRows {
			return errors.New("task not found")
		}

		// If it's a different error (like a dropped connection), return that
	}

	return nil

}

func (t Task) SaveTask() error {
	query := "INSERT INTO tasks (title, IsFinished) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.Title, t.IsFinished)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	t.Id = id
	return nil
}

func (t Task) UpdateTask() error {

	query := "UPDATE tasks SET isFinished=?, title=? WHERE id=?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.IsFinished, t.Title, t.Id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (t Task) DeleteTask() error {
	query := "DELETE FROM tasks WHERE id=?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.Id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

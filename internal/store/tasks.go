package store

import (
	"database/sql"
	"errors"
	"time"
)

type Task struct {
	Id         int64
	Title      string `binding:"required"`
	IsFinished bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TaskStore struct {
	db *sql.DB
}

func (t *TaskStore) GetAllTasks() ([]Task, error) {
	query := "SELECT id, title, isFinished, createdAt, updatedAt FROM tasks"

	rows, err := t.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Title, &task.IsFinished, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TaskStore) GetSingleTask(id int64) (*Task, error) {
	query := "SELECT id, title, isFinished, createdAt, updatedAt FROM tasks WHERE id = ?"

	var task Task
	err := t.db.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.IsFinished, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

func (t *TaskStore) SaveTask(title string, isFinished bool) (int64, error) {
	query := "INSERT INTO tasks (title, IsFinished) VALUES (?, ?)"

	stmt, err := t.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(title, isFinished)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskStore) UpdateTask(isFinished bool, title string, id int64) error {

	query := "UPDATE tasks SET isFinished=?, title=? WHERE id=?"

	stmt, err := t.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(isFinished, title, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (t *TaskStore) DeleteTask(id int64) error {
	query := "DELETE FROM tasks WHERE id=?"

	stmt, err := t.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

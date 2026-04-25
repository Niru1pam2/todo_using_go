package store

import "database/sql"

type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	GetSingleTask(id int64) (*Task, error)
	SaveTask(title string, isFinished bool) (int64, error)
	UpdateTask(isFinished bool, title string, id int64) error
	DeleteTask(id int64) error
}

type Storage struct {
	Tasks TaskRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TaskStore{db},
	}
}

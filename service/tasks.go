package service

import (
	"errors"
	"strings"
	"todo_app/internal/store"
)

// TaskService holds the database store
type TaskService struct {
	store store.Storage
}

func NewTaskService(store store.Storage) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) GetAllTasks() ([]store.Task, error) {
	return s.store.Tasks.GetAllTasks()
}

func (s *TaskService) CreateTask(title string, isFinished bool) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("task title cannot be empty")
	}

	if len(title) < 3 {
		return errors.New("task title must be at least 3 characters long")
	}

	_, err := s.store.Tasks.SaveTask(title, isFinished)
	return err
}

func (s *TaskService) UpdateTask(id int64, title string, isFinished bool) error {

	if id <= 0 {
		return errors.New("invalid task ID")
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return errors.New("task title cannot be empty")
	}

	return s.store.Tasks.UpdateTask(isFinished, title, id)
}

func (s *TaskService) DeleteTask(id int64) error {
	if id <= 0 {
		return errors.New("invalid task ID")
	}
	return s.store.Tasks.DeleteTask(id)
}

func (s *TaskService) GetSingleTask(id int64) (*store.Task, error) {
	if id <= 0 {
		return nil, errors.New("invalid task ID")
	}
	return s.store.Tasks.GetSingleTask(id)
}

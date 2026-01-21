package service

import (
	"time"

	"task_management/model"
	"task_management/repository"

	"github.com/google/uuid"
)

type TaskService struct {
	Repo  *repository.TaskRepository
	Queue chan string
}

func (s *TaskService) Create(task *model.Task) error {
	task.ID = uuid.New().String()
	task.Status = "pending"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err := s.Repo.Create(task)
	if err != nil {
		return err
	}

	// send to worker
	s.Queue <- task.ID
	return nil
}

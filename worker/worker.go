package worker

import (
	"errors"
	"strconv"
	"time"

	"task_management/logger"
	"task_management/repository"
)

type Worker struct {
	Repo  *repository.TaskRepository
	Queue chan string
	Delay time.Duration
}

func New(repo *repository.TaskRepository, queue chan string, minutes string) (*Worker, error) {
	m, err := strconv.Atoi(minutes)
	if err != nil || m <= 0 {
		logger.Error("Invalid AUTO_COMPLETE_MINUTES value: %s", minutes)
		return nil, errors.New("invalid AUTO_COMPLETE_MINUTES")
	}
	logger.Info("Worker initialized with auto-complete delay=%d minutes", m)
	return &Worker{
		Repo:  repo,
		Queue: queue,
		Delay: time.Duration(m) * time.Minute,
	}, nil
}

func (w *Worker) Start() {
	logger.Info("Background worker started")

	go func() {
		for taskID := range w.Queue {
			logger.Info("Received task %s for auto-completion", taskID)
			go w.process(taskID)
		}
	}()
}

func (w *Worker) process(id string) {
	time.Sleep(w.Delay)

	// Only update if still pending/in_progress
	w.Repo.AutoComplete(id)
	logger.Info("Task auto-completed successfully: %s", id)
}

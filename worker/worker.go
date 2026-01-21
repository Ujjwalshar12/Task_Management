package worker

import (
	"errors"
	"strconv"
	"time"

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
		return nil, errors.New("invalid AUTO_COMPLETE_MINUTES")
	}

	return &Worker{
		Repo:  repo,
		Queue: queue,
		Delay: time.Duration(m) * time.Minute,
	}, nil
}

func (w *Worker) Start() {
	go func() {
		for taskID := range w.Queue {
			go w.process(taskID)
		}
	}()
}

func (w *Worker) process(id string) {
	time.Sleep(w.Delay)

	// Only update if still pending/in_progress
	w.Repo.AutoComplete(id)
}

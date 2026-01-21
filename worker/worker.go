package worker

import (
	"strconv"
	"time"

	"task_management/repository"
)

type Worker struct {
	Repo  *repository.TaskRepository
	Queue chan string
	Delay time.Duration
}

func New(repo *repository.TaskRepository, queue chan string, minutes string) *Worker {
	m, _ := strconv.Atoi(minutes)
	return &Worker{
		Repo:  repo,
		Queue: queue,
		Delay: time.Duration(m) * time.Minute,
	}
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

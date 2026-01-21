package main

import (
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"task_management/config"
	"task_management/db"
	"task_management/logger"
	"task_management/repository"
	"task_management/routers"
	"task_management/service"
	"task_management/worker"
)

var cfg *config.Config

func main() {
	//load env variable
	cfg = config.Load()
	logger.Info("Configuration initialized in init()")
	// Connect DB
	dbConn := db.Connect(cfg)
	defer dbConn.Close()

	// Initialize repositories
	taskRepo := &repository.TaskRepository{DB: dbConn}
	userRepo := &repository.UserRepository{DB: dbConn}

	logger.Info("Repositories initialized")

	// Initialize queue & background worker
	queue := make(chan string, 100)

	w := worker.New(taskRepo, queue, cfg.AutoMinutes)
	w.Start()

	logger.Info("Background worker started with auto-complete=%d minutes", cfg.AutoMinutes)

	// Initialize services
	taskService := &service.TaskService{
		Repo:  taskRepo,
		Queue: queue,
	}

	logger.Info("Task service initialized")

	// Setup router (delegated to routers package )
	r := routers.SetupRouter(cfg, taskService, taskRepo, userRepo)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	logger.Info("Server starting on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Server failed to start: %v", err)
	}
}

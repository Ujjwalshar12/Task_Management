package main

import (
	"net/http"

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

	// Load configuration
	cfg := config.Load()
	logger.Info("Configuration Loaded")

	// Connect to database
	dbConn := db.Connect(cfg)
	defer dbConn.Close()

	logger.Info("Database connected")

	// Initialize repositories
	taskRepo := &repository.TaskRepository{DB: dbConn}
	userRepo := &repository.UserRepository{DB: dbConn}

	logger.Info("Repositories initialized")

	// Initialize queue & background worker
	queue := make(chan string, 100)

	w, err := worker.New(taskRepo, queue, cfg.AutoMinutes)
	if err != nil {
		logger.Fatal("Worker initialization failed: %v", err)
	}
	w.Start()

	logger.Info("Background worker started with auto-complete=%d minutes", cfg.AutoMinutes)

	// Initialize services
	taskService := &service.TaskService{
		Repo:  taskRepo,
		Queue: queue,
	}

	logger.Info("Task service initialized")

	// Setup router
	r := routers.SetupRouter(cfg, taskService, taskRepo, userRepo)

	// Start HTTP server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	logger.Info("Server starting on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Server failed to start: %v", err)
	}
}

package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"task_management/config"
	"task_management/handler"
	"task_management/logger"
	"task_management/middleware"
	"task_management/repository"
	"task_management/service"
)

func SetupRouter(
	cfg *config.Config,
	taskService *service.TaskService,
	taskRepo *repository.TaskRepository,
	userRepo *repository.UserRepository,
) *mux.Router {

	r := mux.NewRouter()

	// 🔓 Public routes
	r.HandleFunc("/signup", handler.Signup(userRepo)).Methods("POST")
	r.HandleFunc("/login", handler.Login(userRepo)).Methods("POST")

	logger.Info("Public routes registered")

	// 🔐 Protected routes (JWT middleware)
	auth := middleware.AuthMiddleware(cfg.JWTSecret)

	r.Handle("/tasks", auth(http.HandlerFunc(handler.CreateTask(taskService)))).Methods("POST")
	r.Handle("/tasks", auth(http.HandlerFunc(handler.ListTasks(taskRepo)))).Methods("GET")
	r.Handle("/tasks/{id}", auth(http.HandlerFunc(handler.GetTask(taskRepo)))).Methods("GET")
	r.Handle("/tasks/{id}", auth(http.HandlerFunc(handler.DeleteTask(taskRepo)))).Methods("DELETE")

	logger.Info("Protected routes registered")

	return r
}

package handler

import (
	"encoding/json"
	"net/http"

	"task_management/repository"
)

func ListTasks(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value("user_id").(string)
		role := r.Context().Value("role").(string)

		tasks, err := repo.List(userID, role)
		if err != nil {
			http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

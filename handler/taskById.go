package handler

import (
	"encoding/json"
	"net/http"
	"task_management/repository"

	"github.com/gorilla/mux"
)

func GetTask(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value("user_id").(string)
		role := r.Context().Value("role").(string)

		id := mux.Vars(r)["id"]

		task, err := repo.GetByID(id, userID, role)
		if err != nil {
			http.Error(w, "task not found or access denied", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(task)
	}
}

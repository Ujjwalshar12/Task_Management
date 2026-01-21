package handler

import (
	"net/http"
	"task_management/repository"

	"github.com/gorilla/mux"
)

func DeleteTask(repo *repository.TaskRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value("user_id").(string)
		role := r.Context().Value("role").(string)

		id := mux.Vars(r)["id"]

		err := repo.Delete(id, userID, role)
		if err != nil {
			http.Error(w, "failed to delete task or access denied", http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

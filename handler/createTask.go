package handler

import (
	"encoding/json"
	"net/http"
	"task_management/model"
	"task_management/service"
)

func CreateTask(s *service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value("user_id").(string)

		var input model.Task
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// validation
		if input.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		input.UserID = userID

		err := s.Create(&input)
		if err != nil {
			http.Error(w, "failed to create task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(input)
	}
}

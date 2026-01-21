package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"task_management/middleware"
	"task_management/model"
	"task_management/repository"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Signup(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", 400)
			return
		}

		// validation
		if req.Email == "" || req.Password == "" {
			http.Error(w, "email and password required", 400)
			return
		}

		// default role
		if req.Role == "" {
			req.Role = "user"
		}

		// hash password
		hash, err := middleware.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "password hashing failed", 500)
			return
		}

		user := &model.User{
			ID:       uuid.New().String(),
			Email:    req.Email,
			Password: hash,
			Role:     req.Role,
		}

		err = userRepo.Create(user)
		if err != nil {
			http.Error(w, "email already exists", http.StatusConflict)
			return
		}

		// response without password
		json.NewEncoder(w).Encode(map[string]string{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		})
	}
}

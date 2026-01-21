package handler

import (
	"encoding/json"
	"net/http"

	"task_management/middleware"
	"task_management/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", 400)
			return
		}

		user, err := userRepo.GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", 401)
			return
		}

		err = middleware.CheckPassword(user.Password, req.Password)
		if err != nil {
			http.Error(w, "invalid credentials", 401)
			return
		}

		token, err := middleware.GenerateJWT(*user)
		if err != nil {
			http.Error(w, "token error", 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}

package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"task_management/logger"
	"task_management/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user model.User) (string, error) {

	logger.Info("Generating JWT for user_id=%s role=%s", user.ID, user.Role)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Error("JWT_SECRET is missing in environment")
		return "", fmt.Errorf("jwt secret not configured")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Error("Failed to sign JWT for user_id=%s err=%v", user.ID, err)
		return "", err
	}

	logger.Info("JWT generated successfully for user_id=%s", user.ID)

	return signedToken, nil
}

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			logger.Info("Auth middleware triggered for %s %s", r.Method, r.URL.Path)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Error("Missing Authorization header")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader {
				logger.Error("Invalid Authorization format, expected Bearer token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				logger.Error("JWT parse error: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				logger.Error("Invalid JWT token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				logger.Error("Invalid JWT claims type")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID := claims["user_id"]
			role := claims["role"]

			logger.Info("Authenticated user_id=%v role=%v", userID, role)

			ctx := context.WithValue(r.Context(), "user_id", userID)
			ctx = context.WithValue(ctx, "role", role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

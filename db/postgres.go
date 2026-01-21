package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"task_management/config"
	"task_management/logger"
)

func Connect(cfg *config.Config) *sql.DB {

	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	logger.Info("Connecting to database...")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Fatal("DB connection failed: %v", err)
	}

	if err := db.Ping(); err != nil {
		logger.Fatal("DB not reachable: %v", err)
	}

	logger.Info("Database connection established")

	return db
}

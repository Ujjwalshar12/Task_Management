package logger

import (
	"log"
	"time"
)

func format(level string, msg string, args ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	prefix := "[" + timestamp + "] " + level + " | "
	log.Printf(prefix+msg, args...)
}

// Info level
func Info(msg string, args ...any) {
	format("INFO ", msg, args...)
}

// Error level
func Error(msg string, args ...any) {
	format("ERROR", msg, args...)
}

// Fatal level (exits program)
func Fatal(msg string, args ...any) {
	format("FATAL", msg, args...)
	log.Fatal() // exits
}

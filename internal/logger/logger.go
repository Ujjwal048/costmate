package logger

import (
	"log"
	"os"
	"time"
)

var (
	// Logger is the global logger instance
	Logger *log.Logger
)

// Initialize sets up the logger to write to file only
func Initialize() error {
	// Open log file
	logFile, err := os.OpenFile("/tmp/costmate-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Initialize global logger to write only to file
	Logger = log.New(logFile, "", log.LstdFlags)

	// Log initialization
	Logger.Printf("Logger initialized at %s", time.Now().Format(time.RFC3339))

	return nil
}

// Close closes the log file
func Close() error {
	// The log file is managed by the application lifecycle
	// No need to explicitly close it here
	return nil
}

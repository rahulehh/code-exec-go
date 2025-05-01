package utils

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func init() {
	file, err := os.OpenFile("logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(
		file,
		"message:",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.Lmsgprefix,
	)
}

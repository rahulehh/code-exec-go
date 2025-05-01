package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/rahulehh/code-exec-go/internals/handler"
	"github.com/rahulehh/code-exec-go/internals/utils"
)

func main() {
	socketPath := "/tmp/code-exec.sock"

	if _, err := os.Stat(socketPath); err == nil {
		if err := os.Remove(socketPath); err != nil {
			utils.Logger.Fatalf("Failed to remove existing socket: %v", err)
		}
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		utils.Logger.Fatalf("Failed to listen on unix socket: %v", err)
	}

	if err := os.Chmod(socketPath, 0600); err != nil {
		utils.Logger.Fatalf("Failed to set permissions on socket: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleCodeExecution)

	server := &http.Server{
		Handler: mux,
	}

	fmt.Printf("Server running on Unix socket %s\n", socketPath)
	utils.Logger.Fatal(server.Serve(listener))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/rahulehh/code-exec-go/internals/handler"
	"github.com/rahulehh/code-exec-go/internals/utils"
)

func main() {
	path := ":4202"
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleCodeExecution)

	server := &http.Server{
		Addr:    path,
		Handler: mux,
	}

	fmt.Printf("Server Listening at 127.0.0.1%s\n", path)
	utils.Logger.Fatal(server.ListenAndServe())
}

package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rahulehh/code-exec-go/internals/executor"
	"github.com/rahulehh/code-exec-go/internals/models"
	"github.com/rahulehh/code-exec-go/internals/utils"
)

func HandleCodeExecution(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body:"+err.Error(), http.StatusBadRequest)
		return
	}

	var execRequest models.ExecuteRequest
	err = json.Unmarshal(body, &execRequest)
	if err != nil {
		http.Error(w, "Wrong body format:"+err.Error(), http.StatusBadRequest)
		return
	}

	utils.Logger.Printf("Received data - %s", func() []byte {
		data, err := json.Marshal(execRequest)
		if err != nil {
			utils.Logger.Fatalln("Error request body")
		}
		return data
	}())

	execResult, err := executor.Execute(execRequest)
	if err != nil {
		http.Error(w, "Internal Server Error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(execResult)
}

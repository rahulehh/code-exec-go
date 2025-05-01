package executor

import (
	"github.com/rahulehh/code-exec-go/internals/models"
	"github.com/rahulehh/code-exec-go/internals/utils"
)

// executeInContainer runs the given code snippet inside a language-specific container.
// This assumes the container is already running and accessible via Docker.
func runInContainer(code, language string) models.ExecuteResponse {
	utils.IgnoreUnusedVariables(code, language)
	return models.ExecuteResponse{
		Stdout: "Hello World\n",
		Error:  "",
	}
}

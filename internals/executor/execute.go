package executor

import "github.com/rahulehh/code-exec-go/internals/models"

// Execute handles incoming code execution requests. It ensures a language-specific
// container is running, and then delegates code execution to the appropriate runtime.
func Execute(execRequest models.ExecuteRequest) (models.ExecuteResponse, error) {
	if !isLanguageSupported(execRequest.Language) {
		return models.ExecuteResponse{Stdout: "", Error: "Language not Supported"}, nil
	}

	if !isContainerActive(execRequest.Language) {
		go manageContainerLifecycle(execRequest.Language)
	}

	return runInContainer(execRequest.Code, execRequest.Language), nil
}

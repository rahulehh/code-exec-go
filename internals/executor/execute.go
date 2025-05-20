package executor

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/rahulehh/code-exec-go/internals/models"
	"github.com/rahulehh/code-exec-go/internals/utils"
)

// Execute handles incoming code execution requests. It ensures a language-specific
// container is running, and then delegates code execution to the appropriate runtime.
func Execute(execRequest models.ExecuteRequest) (models.ExecuteResponse, error) {
	if !isLanguageSupported(execRequest.Language) {
		return models.ExecuteResponse{Stdout: "", Error: "Language not Supported"}, nil
	}
	var stdoutBuf, stderrBuf bytes.Buffer

	var image string
	switch execRequest.Language {
	case "python":
		image = "python-runner"
	case "go":
		image = "go-runner"
	case "c":
		image = "c-runner"
	case "cpp":
		image = "cpp-runner"
	default:
		return models.ExecuteResponse{
			Error: "Unsupported language: " + execRequest.Language,
		}, nil
	}

	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-i",
		"--network=none",
		"--cpus=2",
		"--memory=512m",
		"--pids-limit=50",
		"--security-opt=no-new-privileges",
		"--cap-drop=ALL",
		image,
	)
	cmd.Stdin = strings.NewReader(execRequest.Code)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	if err != nil {
		utils.Logger.Printf(
			"Error running the program\nimage: %s\ncode:\n%s",
			image,
			execRequest.Code,
		)
		return models.ExecuteResponse{
			Stdout: stdoutBuf.String(),
			Stderr: stderrBuf.String(),
			Error:  "Error occured in server",
		}, nil
	}

	return models.ExecuteResponse{
		Stdout: stdoutBuf.String(),
		Stderr: stderrBuf.String(),
	}, nil
}

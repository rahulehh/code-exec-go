package executor_test

import (
	"testing"

	"github.com/rahulehh/code-exec-go/internals/executor"
	"github.com/rahulehh/code-exec-go/internals/models"
)

func TestExecute_Python(t *testing.T) {
	output, err := executor.Execute(models.ExecuteRequest{
		Language: "python",
		Code:     "print('Hello World')",
	})
	if err != nil {
		t.Fatalf("Unexpected error from Execute: %v", err)
	}
	if output.Error != "" {
		t.Fatalf("Program execution returned error: %s", output.Error)
	}
	expectedOutput := "Hello World\n"
	if output.Stdout != expectedOutput {
		t.Fatalf("Unexpected output:\nexpected: %q\ngot: %q", expectedOutput, output.Stdout)
	}
}

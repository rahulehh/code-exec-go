package executor_test

import (
	"testing"

	"github.com/rahulehh/code-exec-go/internals/executor"
	"github.com/rahulehh/code-exec-go/internals/models"
)

func TestExecute_Python(t *testing.T) {
	output, err := executor.Execute(models.ExecuteRequest{
		Language: "python",
		Code:     "print(\"Hello from Python\")",
	})
	if err != nil {
		t.Fatalf("Unexpected error from Execute: %v", err)
	}
	if output.Error != "" {
		t.Fatalf("Program execution returned error: %s", output.Error)
	}
	expectedOutput := "Hello from Python\n"
	if output.Stdout != expectedOutput {
		t.Fatalf("Unexpected output:\nexpected: %q\ngot: %q", expectedOutput, output.Stdout)
	}
}

func TestExecute_Go(t *testing.T) {
	output, err := executor.Execute(models.ExecuteRequest{
		Language: "python",
		Code:     "print(\"Hello from Go\")",
	})
	if err != nil {
		t.Fatalf("Unexpected error from Execute: %v", err)
	}
	if output.Error != "" {
		t.Fatalf("Program execution returned error: %s", output.Error)
	}
	expectedOutput := "Hello from Go\n"
	if output.Stdout != expectedOutput {
		t.Fatalf("Unexpected output:\nexpected: %q\ngot: %q", expectedOutput, output.Stdout)
	}
}

func TestExecute_C(t *testing.T) {
	output, err := executor.Execute(models.ExecuteRequest{
		Language: "python",
		Code:     "print(\"Hello from C\")",
	})
	if err != nil {
		t.Fatalf("Unexpected error from Execute: %v", err)
	}
	if output.Error != "" {
		t.Fatalf("Program execution returned error: %s", output.Error)
	}
	expectedOutput := "Hello from C\n"
	if output.Stdout != expectedOutput {
		t.Fatalf("Unexpected output:\nexpected: %q\ngot: %q", expectedOutput, output.Stdout)
	}
}

func TestExecute_Cpp(t *testing.T) {
	output, err := executor.Execute(models.ExecuteRequest{
		Language: "python",
		Code:     "print(\"Hello from C++\")",
	})
	if err != nil {
		t.Fatalf("Unexpected error from Execute: %v", err)
	}
	if output.Error != "" {
		t.Fatalf("Program execution returned error: %s", output.Error)
	}
	expectedOutput := "Hello from C++\n"
	if output.Stdout != expectedOutput {
		t.Fatalf("Unexpected output:\nexpected: %q\ngot: %q", expectedOutput, output.Stdout)
	}
}

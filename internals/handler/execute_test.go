package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rahulehh/code-exec-go/internals/handler"
	"github.com/rahulehh/code-exec-go/internals/models"
)

func TestHandleCodeExecution_Success(t *testing.T) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(models.ExecuteRequest{
		Language: "python",
		Code:     "print('Hello World')",
	})

	req := httptest.NewRequest(http.MethodPost, "/", &buf)
	w := httptest.NewRecorder()

	handler.HandleCodeExecution(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var respData models.ExecuteResponse
	err := json.NewDecoder(resp.Body).Decode(&respData)

	if err != nil {
		t.Fatalf("response body has different structure")
	}
}

func TestHandleCodeExecution_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.HandleCodeExecution(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status code: %d, got %d",
			http.StatusMethodNotAllowed,
			resp.StatusCode,
		)
	}
}

func TestHandleCodeExecution_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{invalid json`)))
	w := httptest.NewRecorder()

	handler.HandleCodeExecution(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", w.Code)
	}
}

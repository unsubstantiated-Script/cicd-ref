package tests

import (
	"CICDRef/internal/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	h := handlers.New("test")
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	h.Health(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Health check failed")
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("expected status to be ok, got %s", body["status"])
	}
}

func TestMessageHandler(t *testing.T) {
	h := handlers.New("hello pipeline")
	req := httptest.NewRequest(http.MethodGet, "/api/v1/message", nil)
	rr := httptest.NewRecorder()

	h.MessageHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", rr.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "hello pipeline" {
		t.Fatalf("Expected message to be hello pipeline, got %s", body["message"])
	}

}

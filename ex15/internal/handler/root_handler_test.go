package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"takaya-47/learn-go/ex15/internal/handler"
	"testing"
)

func TestNewHandler(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("hello"))
	w := httptest.NewRecorder()
	out := make(chan []byte, 1)

	// Act
	handler.NewHandler(out).ServeHTTP(w, r)

	// Assert
	got := string(<-out)
	if got != "hello" {
		t.Errorf("expected hello, but got %s", got)
	}

	if w.Code != http.StatusAccepted {
		t.Errorf("expected status code %d, but got %d", http.StatusAccepted, w.Code)
	}

	got = w.Body.String()
	if got != "OK: 1" {
		t.Errorf("expected body OK: 1, but got %s", got)
	}
}

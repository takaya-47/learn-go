package handler_test

import (
	"errors"
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
	if got, want := string(<-out), "hello"; got != want {
		t.Errorf("expected channel value is %q, got %q", want, got)
	}

	if got, want := w.Code, http.StatusAccepted; got != want {
		t.Errorf("expected status code is %d, got %d", want, got)
	}

	if got, want := w.Body.String(), "OK: 1"; got != want {
		t.Errorf("expected body is %q, but got %q", want, got)
	}
}

type errorReader struct{}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func TestNewHandler_ReadError(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodPost, "/", errorReader{})
	w := httptest.NewRecorder()
	out := make(chan []byte, 1)

	// Act
	handler.NewHandler(out).ServeHTTP(w, r)

	// Assert
	if got, want := w.Code, http.StatusBadRequest; got != want {
		t.Errorf("expected status code is %d, got %d", want, got)
	}

	if got, want := w.Body.String(), "Bad Input"; got != want {
		t.Errorf("expected body is %q, but got %q", want, got)
	}
}

func TestNewHandler_TooBusy(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("hello"))
	w := httptest.NewRecorder()
	// バッファ無しチャネルのため、NewHandler側で書き込みがブロックされ、defaultケースに入る
	out := make(chan []byte)

	// Act
	handler.NewHandler(out).ServeHTTP(w, r)

	// Assert
	if got, want := w.Code, http.StatusServiceUnavailable; got != want {
		t.Errorf("expected status code is %d, got %d", want, got)
	}

	if got, want := w.Body.String(), "Too Busy: 1"; got != want {
		t.Errorf("expected body is %q, but got %q", want, got)
	}
}

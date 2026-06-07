package http

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func assertStatus(t *testing.T, rec *httptest.ResponseRecorder, want int) {
	t.Helper()

	if rec.Code != want {
		t.Fatalf("status = %d, want %d", rec.Code, want)
	}
}

func assertHTTPError(t *testing.T, err error, want int) {
	t.Helper()

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != want {
		t.Fatalf("status = %d, want %d", httpErr.Code, want)
	}
}

func assertEqual[T comparable](t *testing.T, got T, want T, label string) {
	t.Helper()

	if got != want {
		t.Fatalf("%s = %v, want %v", label, got, want)
	}
}

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestGoogleHandlerStart_UsesConfiguredRedirectWhenQueryMissing(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/auth/google/start?purpose=login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	flow := &fakeGoogleFlow{
		startOutput: authusecase.GoogleStartOutput{
			RedirectTo: "https://example.com/oauth",
			State:      "state-123",
		},
	}

	handler := NewGoogleHandler(flow, "http://127.0.0.1:8081/api/auth/google/callback")

	if err := handler.Start(c); err != nil {
		t.Fatalf("handler.Start() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if flow.startInput.Purpose != "login" {
		t.Fatalf("purpose = %q, want login", flow.startInput.Purpose)
	}
	if flow.startInput.RedirectURL != "http://127.0.0.1:8081/api/auth/google/callback" {
		t.Fatalf("redirect url = %q", flow.startInput.RedirectURL)
	}

	var body authusecase.GoogleStartOutput
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if body.RedirectTo != "https://example.com/oauth" {
		t.Fatalf("redirect_to = %q", body.RedirectTo)
	}
	if body.State != "state-123" {
		t.Fatalf("state = %q", body.State)
	}
}

func TestGoogleHandlerStart_AllowsRedirectOverrideFromQuery(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/auth/google/start?purpose=login&redirect_url=http://127.0.0.1:3000/callback",
		nil,
	)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	flow := &fakeGoogleFlow{
		startOutput: authusecase.GoogleStartOutput{
			RedirectTo: "https://example.com/oauth",
			State:      "state-456",
		},
	}

	handler := NewGoogleHandler(flow, "http://127.0.0.1:8081/api/auth/google/callback")

	if err := handler.Start(c); err != nil {
		t.Fatalf("handler.Start() error = %v", err)
	}

	if flow.startInput.RedirectURL != "http://127.0.0.1:3000/callback" {
		t.Fatalf("redirect url = %q", flow.startInput.RedirectURL)
	}
}

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestGoogleHandlerCallback_UsesConfiguredRedirectURL(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/auth/google/callback?code=code-123&state=state-123&redirect_url=http://evil.local/callback",
		nil,
	)
	req.Header.Set("User-Agent", "handler-test-agent")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	flow := &fakeGoogleFlow{
		callbackOut: authusecase.GoogleCallbackOutput{
			AccessToken:    "access-token",
			AccessExp:      time.Unix(1776647989, 0),
			RefreshToken:   "refresh-token",
			RefreshExp:     time.Unix(1779240089, 0),
			TrustLevel:     "aal1",
			StepUpRequired: false,
		},
	}

	handler := NewGoogleHandler(flow, "http://127.0.0.1:8081/api/auth/google/callback")

	if err := handler.Callback(c); err != nil {
		t.Fatalf("handler.Callback() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if flow.callbackInput.Code != "code-123" {
		t.Fatalf("code = %q", flow.callbackInput.Code)
	}
	if flow.callbackInput.State != "state-123" {
		t.Fatalf("state = %q", flow.callbackInput.State)
	}
	if flow.callbackInput.RedirectURL != "http://127.0.0.1:8081/api/auth/google/callback" {
		t.Fatalf("redirect url = %q", flow.callbackInput.RedirectURL)
	}
	if flow.callbackInput.Client.UserAgent != "handler-test-agent" {
		t.Fatalf("user agent = %q", flow.callbackInput.Client.UserAgent)
	}

	var body authusecase.GoogleCallbackOutput
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if body.AccessToken != "access-token" {
		t.Fatalf("access token = %q", body.AccessToken)
	}
	if body.RefreshToken != "refresh-token" {
		t.Fatalf("refresh token = %q", body.RefreshToken)
	}
	if body.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q", body.TrustLevel)
	}
}

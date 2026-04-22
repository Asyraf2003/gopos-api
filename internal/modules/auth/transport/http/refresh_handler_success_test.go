package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestRefreshHandler_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/refresh", strings.NewReader(`{"refresh_token":"refresh-token-123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	usecase := &fakeRefreshTokenUsecase{
		output: authusecase.RefreshTokenOutput{
			AccessToken:    "new-access-token",
			AccessExp:      time.Unix(1776685000, 0),
			RefreshToken:   "new-refresh-token",
			RefreshExp:     time.Unix(1779277000, 0),
			TrustLevel:     "aal1",
			StepUpRequired: false,
		},
	}

	handler := NewRefreshHandler(usecase)

	if err := handler.Refresh(c); err != nil {
		t.Fatalf("Refresh() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if usecase.lastInput.RefreshToken != "refresh-token-123" {
		t.Fatalf("refresh token input = %q", usecase.lastInput.RefreshToken)
	}

	var body authusecase.RefreshTokenOutput
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if body.AccessToken != "new-access-token" {
		t.Fatalf("access token = %q", body.AccessToken)
	}
	if body.RefreshToken != "new-refresh-token" {
		t.Fatalf("refresh token = %q", body.RefreshToken)
	}
	if body.TrustLevel != "aal1" {
		t.Fatalf("trust level = %q", body.TrustLevel)
	}
}

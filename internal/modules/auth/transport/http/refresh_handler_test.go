package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type fakeRefreshTokenUsecase struct {
	lastInput authusecase.RefreshTokenInput
	output    authusecase.RefreshTokenOutput
	err       error
}

func (f *fakeRefreshTokenUsecase) Execute(ctx context.Context, in authusecase.RefreshTokenInput) (authusecase.RefreshTokenOutput, error) {
	_ = ctx
	f.lastInput = in
	return f.output, f.err
}

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

func TestRefreshHandler_RejectsInvalidBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/refresh", strings.NewReader(`{"refresh_token":`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewRefreshHandler(&fakeRefreshTokenUsecase{})

	err := handler.Refresh(c)
	if err == nil {
		t.Fatal("Refresh() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
}

func TestRefreshHandler_RejectsInvalidRefreshToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/refresh", strings.NewReader(`{"refresh_token":"old-token"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewRefreshHandler(&fakeRefreshTokenUsecase{
		err: authusecase.ErrInvalidRefreshToken,
	})

	err := handler.Refresh(c)
	if err == nil {
		t.Fatal("Refresh() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", httpErr.Code)
	}
}

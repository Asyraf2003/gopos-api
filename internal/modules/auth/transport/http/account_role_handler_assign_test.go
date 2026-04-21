package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAccountRoleHandler_AssignSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/accounts/acc-123/roles", strings.NewReader(`{"role_key":"admin"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("account_id")
	c.SetParamValues("acc-123")

	handler, assignUsecase, _ := newAccountRoleHandlerForTest(nil, nil)

	if err := handler.Assign(c); err != nil {
		t.Fatalf("Assign() error = %v", err)
	}

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}
	if assignUsecase.calls != 1 {
		t.Fatalf("assign calls = %d, want 1", assignUsecase.calls)
	}
	if assignUsecase.lastAccountID != "acc-123" {
		t.Fatalf("account id = %q", assignUsecase.lastAccountID)
	}
	if assignUsecase.lastRoleKey != "admin" {
		t.Fatalf("role key = %q", assignUsecase.lastRoleKey)
	}
}

func TestAccountRoleHandler_AssignRejectsInvalidBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/accounts/acc-123/roles", strings.NewReader(`{"role_key":`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("account_id")
	c.SetParamValues("acc-123")

	handler, _, _ := newAccountRoleHandlerForTest(nil, nil)

	err := handler.Assign(c)
	if err == nil {
		t.Fatal("Assign() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
}

func TestAccountRoleHandler_AssignRejectsMissingRoleKey(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/accounts/acc-123/roles", strings.NewReader(`{"role_key":"   "}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("account_id")
	c.SetParamValues("acc-123")

	handler, assignUsecase, _ := newAccountRoleHandlerForTest(nil, nil)

	err := handler.Assign(c)
	if err == nil {
		t.Fatal("Assign() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
	if assignUsecase.calls != 0 {
		t.Fatalf("assign calls = %d, want 0", assignUsecase.calls)
	}
}

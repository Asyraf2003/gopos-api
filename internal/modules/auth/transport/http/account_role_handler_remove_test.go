package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

func TestAccountRoleHandler_RemoveSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/admin/accounts/acc-123/roles/admin", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("account_id", "role_key")
	c.SetParamValues("acc-123", "admin")

	handler, _, removeUsecase := newAccountRoleHandlerForTest(nil, nil)

	if err := handler.Remove(c); err != nil {
		t.Fatalf("Remove() error = %v", err)
	}

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}
	if removeUsecase.calls != 1 {
		t.Fatalf("remove calls = %d, want 1", removeUsecase.calls)
	}
	if removeUsecase.lastAccountID != "acc-123" {
		t.Fatalf("account id = %q", removeUsecase.lastAccountID)
	}
	if removeUsecase.lastRoleKey != "admin" {
		t.Fatalf("role key = %q", removeUsecase.lastRoleKey)
	}
}

func TestAccountRoleHandler_RemoveRejectsBaseRole(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/admin/accounts/acc-123/roles/base", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("account_id", "role_key")
	c.SetParamValues("acc-123", "base")

	handler, _, _ := newAccountRoleHandlerForTest(nil, authusecase.ErrBaseRoleRemovalNotAllowed)

	err := handler.Remove(c)
	if err == nil {
		t.Fatal("Remove() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/auth/domain"

	"github.com/labstack/echo/v4"
)

func TestRequirePermission_AllowsWhenPermissionExists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(WithPrincipal(req.Context(), domain.Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles:     []string{"cashier"},
		Permissions: []string{
			"sale.order.create",
			"sale.order.read",
		},
		TrustLevel: "aal1",
	}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	called := false
	handler := RequirePermission("sale.order.create")(func(c echo.Context) error {
		called = true
		return c.NoContent(http.StatusNoContent)
	})

	if err := handler(c); err != nil {
		t.Fatalf("handler() error = %v", err)
	}

	if !called {
		t.Fatal("next handler was not called")
	}
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}
}

func TestRequirePermission_RejectsMissingPrincipal(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RequirePermission("sale.order.create")(func(c echo.Context) error {
		t.Fatal("next handler should not be called")
		return nil
	})

	err := handler(c)
	if err == nil {
		t.Fatal("handler() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", httpErr.Code)
	}
}

func TestRequirePermission_RejectsMissingPermission(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(WithPrincipal(req.Context(), domain.Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles:     []string{"base"},
		Permissions: []string{
			"profile.self.read",
		},
		TrustLevel: "aal1",
	}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RequirePermission("inventory.manage")(func(c echo.Context) error {
		t.Fatal("next handler should not be called")
		return nil
	})

	err := handler(c)
	if err == nil {
		t.Fatal("handler() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", httpErr.Code)
	}
}

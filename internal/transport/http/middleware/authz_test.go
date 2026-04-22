package middleware

import (
	"net/http"
	"testing"

	"pos-go/internal/modules/auth/domain"

	"github.com/labstack/echo/v4"
)

func TestRequirePermission_AllowsWhenPermissionExists(t *testing.T) {
	c, rec := newAuthzTestContext(&domain.Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles:     []string{"cashier"},
		Permissions: []string{
			"sale.order.create",
			"sale.order.read",
		},
		TrustLevel: "aal1",
	})

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
	c, _ := newAuthzTestContext(nil)

	handler := RequirePermission("sale.order.create")(func(c echo.Context) error {
		t.Fatal("next handler should not be called")
		return nil
	})

	assertHTTPErrorCode(t, handler(c), http.StatusUnauthorized)
}

func TestRequirePermission_RejectsMissingPermission(t *testing.T) {
	c, _ := newAuthzTestContext(&domain.Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles:     []string{"base"},
		Permissions: []string{
			"profile.self.read",
		},
		TrustLevel: "aal1",
	})

	handler := RequirePermission("inventory.manage")(func(c echo.Context) error {
		t.Fatal("next handler should not be called")
		return nil
	})

	assertHTTPErrorCode(t, handler(c), http.StatusForbidden)
}

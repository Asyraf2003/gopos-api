package middleware

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"pos-go/internal/modules/capability/domain"

	"github.com/labstack/echo/v4"
)

func TestRequireCapability_AllowsEnabledCapability(t *testing.T) {
	c, rec := newCapabilityTestContext()
	checker := capabilityCheckerFunc(func(ctx context.Context, key string) error {
		_ = ctx
		if key != "account.role.assign" {
			t.Fatalf("capability key = %q", key)
		}
		return nil
	})

	called := false
	handler := RequireCapability("account.role.assign", checker)(func(c echo.Context) error {
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

func TestRequireCapability_RejectsDisabledBeforeHandler(t *testing.T) {
	c, _ := newCapabilityTestContext()
	checker := capabilityCheckerFunc(func(ctx context.Context, key string) error {
		_ = ctx
		_ = key
		return domain.ErrCapabilityDisabled
	})

	handler := RequireCapability("account.role.assign", checker)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusForbidden)
}

func TestRequireCapability_RejectsCheckerErrorBeforeHandler(t *testing.T) {
	c, _ := newCapabilityTestContext()
	checker := capabilityCheckerFunc(func(ctx context.Context, key string) error {
		_ = ctx
		_ = key
		return errors.New("repository unavailable")
	})

	handler := RequireCapability("account.role.assign", checker)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusInternalServerError)
}

func TestRequireCapability_RejectsEmptyKey(t *testing.T) {
	c, _ := newCapabilityTestContext()

	handler := RequireCapability(" ", capabilityCheckerFunc(func(ctx context.Context, key string) error {
		t.Fatal("checker should not be called")
		return nil
	}))(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusInternalServerError)
}

func TestRequireCapability_RejectsNilChecker(t *testing.T) {
	c, _ := newCapabilityTestContext()

	handler := RequireCapability("account.role.assign", nil)(failIfCalledHandler(t))

	assertHTTPErrorCode(t, handler(c), http.StatusInternalServerError)
}

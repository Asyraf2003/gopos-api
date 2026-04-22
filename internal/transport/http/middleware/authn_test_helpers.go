package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/auth/ports"

	"github.com/labstack/echo/v4"
)

func newAuthnTestContext(authHeader string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if authHeader != "" {
		req.Header.Set(echo.HeaderAuthorization, authHeader)
	}

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func failIfCalledHandler(t *testing.T) echo.HandlerFunc {
	t.Helper()

	return func(c echo.Context) error {
		t.Fatal("next handler should not be called")
		return nil
	}
}

func assertPrincipalInContext(t *testing.T, ctx context.Context, wantAccountID string, wantPermission string) {
	t.Helper()

	principal, ok := PrincipalFromContext(ctx)
	if !ok {
		t.Fatal("principal missing from context")
	}
	if principal.AccountID != wantAccountID {
		t.Fatalf("account id = %q", principal.AccountID)
	}
	if !principal.HasPermission(wantPermission) {
		t.Fatalf("expected permission %s", wantPermission)
	}
}

func assertResolverInput(
	t *testing.T,
	got ports.ResolvePrincipalInput,
	wantAccountID string,
	wantSessionID string,
	wantTrustLevel string,
) {
	t.Helper()

	if got.AccountID != wantAccountID {
		t.Fatalf("resolver account id = %q", got.AccountID)
	}
	if got.SessionID != wantSessionID {
		t.Fatalf("resolver session id = %q", got.SessionID)
	}
	if got.TrustLevel != wantTrustLevel {
		t.Fatalf("resolver trust level = %q", got.TrustLevel)
	}
}

func assertSessionCheckerSessionID(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Fatalf("session checker session id = %q", got)
	}
}

package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/labstack/echo/v4"
)

type fakeAccessTokenVerifier struct {
	claims ports.AccessTokenClaims
	err    error
}

func (f *fakeAccessTokenVerifier) VerifyAccessToken(ctx context.Context, token string) (ports.AccessTokenClaims, error) {
	_ = ctx
	_ = token
	return f.claims, f.err
}

type fakePrincipalResolver struct {
	principal domain.Principal
	err       error
	lastInput ports.ResolvePrincipalInput
}

func (f *fakePrincipalResolver) Resolve(ctx context.Context, in ports.ResolvePrincipalInput) (domain.Principal, error) {
	_ = ctx
	f.lastInput = in
	return f.principal, f.err
}

type fakeSessionStatusChecker struct {
	active        bool
	err           error
	lastSessionID string
}

func (f *fakeSessionStatusChecker) IsSessionActive(ctx context.Context, sessionID string) (bool, error) {
	_ = ctx
	f.lastSessionID = sessionID
	return f.active, f.err
}

func TestRequireAuth_SetsPrincipalOnSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer token-123")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	verifier := &fakeAccessTokenVerifier{
		claims: ports.AccessTokenClaims{
			AccountID:  "acc-123",
			SessionID:  "sess-123",
			TrustLevel: "aal1",
		},
	}
	resolver := &fakePrincipalResolver{
		principal: domain.Principal{
			AccountID:   "acc-123",
			SessionID:   "sess-123",
			Roles:       []string{"base"},
			Permissions: []string{"profile.self.read"},
			TrustLevel:  "aal1",
		},
	}
	sessionChecker := &fakeSessionStatusChecker{active: true}

	called := false
	handler := RequireAuth(verifier, resolver, sessionChecker)(func(c echo.Context) error {
		called = true

		principal, ok := PrincipalFromContext(c.Request().Context())
		if !ok {
			t.Fatal("principal missing from context")
		}

		if principal.AccountID != "acc-123" {
			t.Fatalf("account id = %q", principal.AccountID)
		}
		if !principal.HasPermission("profile.self.read") {
			t.Fatal("expected permission profile.self.read")
		}

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

	if resolver.lastInput.AccountID != "acc-123" {
		t.Fatalf("resolver account id = %q", resolver.lastInput.AccountID)
	}
	if resolver.lastInput.SessionID != "sess-123" {
		t.Fatalf("resolver session id = %q", resolver.lastInput.SessionID)
	}
	if resolver.lastInput.TrustLevel != "aal1" {
		t.Fatalf("resolver trust level = %q", resolver.lastInput.TrustLevel)
	}
	if sessionChecker.lastSessionID != "sess-123" {
		t.Fatalf("session checker session id = %q", sessionChecker.lastSessionID)
	}
}

func TestRequireAuth_RejectsMissingBearerToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RequireAuth(
		&fakeAccessTokenVerifier{},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: true},
	)(func(c echo.Context) error {
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

func TestRequireAuth_RejectsInvalidAccessToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer token-123")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RequireAuth(
		&fakeAccessTokenVerifier{err: errors.New("bad token")},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: true},
	)(func(c echo.Context) error {
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

func TestRequireAuth_RejectsInactiveSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer token-123")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := RequireAuth(
		&fakeAccessTokenVerifier{
			claims: ports.AccessTokenClaims{
				AccountID:  "acc-123",
				SessionID:  "sess-123",
				TrustLevel: "aal1",
			},
		},
		&fakePrincipalResolver{},
		&fakeSessionStatusChecker{active: false},
	)(func(c echo.Context) error {
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

package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/capability/domain"

	"github.com/labstack/echo/v4"
)

func TestProtectedRoutesRejectDisabledCapabilityBeforeHandler(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		routePath     string
		requestPath   string
		capabilityKey string
	}{
		{"profile self", http.MethodGet, "/api/me", "/api/me", "profile.self.show"},
		{"authz profile self", http.MethodGet, "/api/authz/me", "/api/authz/me", "authz.profile.self.show"},
		{"logout", http.MethodPost, "/api/auth/logout", "/api/auth/logout", "auth.session.logout"},
		{"assign account role", http.MethodPost, "/api/admin/accounts/:account_id/roles", "/api/admin/accounts/acc-1/roles", "account.role.assign"},
		{"remove account role", http.MethodDelete, "/api/admin/accounts/:account_id/roles/:role_key", "/api/admin/accounts/acc-1/roles/admin", "account.role.remove"},
		{"list capabilities", http.MethodGet, "/api/admin/capabilities", "/api/admin/capabilities", "capability.manage"},
		{"show capability", http.MethodGet, "/api/admin/capabilities/:key", "/api/admin/capabilities/profile.self.show", "capability.manage"},
		{"enable capability", http.MethodPost, "/api/admin/capabilities/:key/enable", "/api/admin/capabilities/profile.self.show/enable", "capability.manage"},
		{"disable capability", http.MethodPost, "/api/admin/capabilities/:key/disable", "/api/admin/capabilities/profile.self.show/disable", "capability.manage"},
		{"service catalog list", http.MethodGet, "/api/service-catalog/items", "/api/service-catalog/items", "service_catalog.list"},
		{"service catalog create", http.MethodPost, "/api/service-catalog/items", "/api/service-catalog/items", "service_catalog.create"},
		{"service catalog lookup", http.MethodGet, "/api/service-catalog/items/lookup", "/api/service-catalog/items/lookup", "service_catalog.lookup"},
		{"service catalog show", http.MethodGet, "/api/service-catalog/items/:id", "/api/service-catalog/items/svc_1", "service_catalog.show"},
		{"service catalog update", http.MethodPut, "/api/service-catalog/items/:id", "/api/service-catalog/items/svc_1", "service_catalog.update"},
		{"service catalog activate", http.MethodPost, "/api/service-catalog/items/:id/activate", "/api/service-catalog/items/svc_1/activate", "service_catalog.activate"},
		{"service catalog deactivate", http.MethodPost, "/api/service-catalog/items/:id/deactivate", "/api/service-catalog/items/svc_1/deactivate", "service_catalog.deactivate"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			gotKey := ""

			checker := capabilityCheckerFunc(func(ctx context.Context, key string) error {
				_ = ctx
				gotKey = key
				return domain.ErrCapabilityDisabled
			})

			e.Add(
				tt.method,
				tt.routePath,
				RequireCapability(tt.capabilityKey, checker)(failIfCalledHandler(t)),
			)

			req := httptest.NewRequest(tt.method, tt.requestPath, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			if rec.Code != http.StatusForbidden {
				t.Fatalf("status = %d, want 403", rec.Code)
			}
			if gotKey != tt.capabilityKey {
				t.Fatalf("capability key = %q, want %q", gotKey, tt.capabilityKey)
			}
		})
	}
}

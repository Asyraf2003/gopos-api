package postgres_test

import (
	"context"
	"testing"

	"pos-go/internal/platform/postgres"
)

func TestSeededExistingProtectedCapabilities(t *testing.T) {
	t.Parallel()

	pool := integrationPool(t)
	repo := postgres.NewCapabilityRepository(pool)

	expected := map[string]struct {
		method             string
		path               string
		requiredPermission string
	}{
		"profile.self.show": {
			method:             "GET",
			path:               "/api/me",
			requiredPermission: "profile.self.read",
		},
		"authz.profile.self.show": {
			method:             "GET",
			path:               "/api/authz/me",
			requiredPermission: "profile.self.read",
		},
		"auth.session.logout": {
			method:             "POST",
			path:               "/api/auth/logout",
			requiredPermission: "auth.session.logout",
		},
		"account.role.assign": {
			method:             "POST",
			path:               "/api/admin/accounts/:account_id/roles",
			requiredPermission: "account.role.assign",
		},
		"account.role.remove": {
			method:             "DELETE",
			path:               "/api/admin/accounts/:account_id/roles/:role_key",
			requiredPermission: "account.role.assign",
		},
	}

	for key, want := range expected {
		got, err := repo.Get(context.Background(), key)
		if err != nil {
			t.Fatalf("expected seeded capability %q: %v", key, err)
		}
		if got.Method != want.method {
			t.Fatalf("%s method: got %q want %q", key, got.Method, want.method)
		}
		if got.Path != want.path {
			t.Fatalf("%s path: got %q want %q", key, got.Path, want.path)
		}
		if got.RequiredPermission != want.requiredPermission {
			t.Fatalf("%s permission: got %q want %q", key, got.RequiredPermission, want.requiredPermission)
		}
		if !got.Enabled || !got.DefaultEnabled {
			t.Fatalf("%s should be enabled by default", key)
		}
	}
}

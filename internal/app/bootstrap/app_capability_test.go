package bootstrap

import (
	"os"
	"strings"
	"testing"
)

func TestNew_WiresCapabilityGuardsForProtectedRoutes(t *testing.T) {
	source, err := os.ReadFile("app.go")
	if err != nil {
		t.Fatalf("ReadFile(app.go) error = %v", err)
	}

	required := []string{
		`RequireCapability("profile.self.show"`,
		`RequireCapability("authz.profile.self.show"`,
		`RequireCapability("auth.session.logout"`,
		`RequireCapability("account.role.assign"`,
		`RequireCapability("account.role.remove"`,
		`RequireCapability("capability.manage"`,
	}

	for _, needle := range required {
		if !strings.Contains(string(source), needle) {
			t.Fatalf("missing bootstrap capability guard: %s", needle)
		}
	}
}

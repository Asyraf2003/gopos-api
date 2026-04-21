package domain

import "testing"

func TestPrincipalHasPermission(t *testing.T) {
	principal := Principal{
		AccountID: "acc-123",
		SessionID: "sess-123",
		Roles: []string{
			"base",
			"cashier",
		},
		Permissions: []string{
			"profile.self.read",
			"sale.order.create",
			"sale.order.read",
		},
		TrustLevel: "aal1",
	}

	if !principal.HasPermission("sale.order.create") {
		t.Fatal("expected permission sale.order.create to be present")
	}

	if principal.HasPermission("inventory.manage") {
		t.Fatal("did not expect permission inventory.manage to be present")
	}
}

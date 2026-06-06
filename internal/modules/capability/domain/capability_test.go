package domain

import "testing"

func TestNewCapabilityNormalizesAndValidates(t *testing.T) {
	capability, err := NewCapability(Capability{
		Key:                " account.role.assign ",
		Domain:             " account ",
		Operation:          " assign-role ",
		Method:             " post ",
		Path:               " /api/admin/accounts/:account_id/roles ",
		DefaultEnabled:     true,
		Enabled:            true,
		RequiredPermission: " account.role.assign ",
		RiskLevel:          RiskHigh,
		AuditRequired:      true,
		OwnerPackage:       " internal/modules/auth ",
		TestProof:          " internal/modules/auth/transport/http/account_role_handler_assign_test.go ",
		DisabledReason:     " stale reason ",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	if capability.Key != "account.role.assign" {
		t.Fatalf("key = %q", capability.Key)
	}
	if capability.Method != "POST" {
		t.Fatalf("method = %q", capability.Method)
	}
	if capability.DisabledReason != "" {
		t.Fatalf("disabled reason = %q, want empty", capability.DisabledReason)
	}
}

func TestNewCapabilityRejectsInvalidRisk(t *testing.T) {
	_, err := NewCapability(validCapability(func(c *Capability) {
		c.RiskLevel = "critical"
	}))
	if err == nil {
		t.Fatal("NewCapability() error = nil, want error")
	}
}

func TestCapabilityEnableAndDisable(t *testing.T) {
	capability, err := NewCapability(validCapability(func(c *Capability) {
		c.Enabled = true
	}))
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	disabled := capability.Disable("maintenance")
	if disabled.Enabled {
		t.Fatal("disabled capability is enabled")
	}
	if disabled.DisabledReason != "maintenance" {
		t.Fatalf("disabled reason = %q", disabled.DisabledReason)
	}

	enabled := disabled.Enable()
	if !enabled.Enabled {
		t.Fatal("enabled capability is disabled")
	}
	if enabled.DisabledReason != "" {
		t.Fatalf("enabled disabled reason = %q", enabled.DisabledReason)
	}
}

func validCapability(mutate func(*Capability)) Capability {
	capability := Capability{
		Key:                "account.role.assign",
		Domain:             "account",
		Operation:          "assign-role",
		Method:             "POST",
		Path:               "/api/admin/accounts/:account_id/roles",
		DefaultEnabled:     true,
		Enabled:            true,
		RequiredPermission: "account.role.assign",
		RiskLevel:          RiskHigh,
		AuditRequired:      true,
		OwnerPackage:       "internal/modules/auth",
		TestProof:          "internal/modules/auth/transport/http/account_role_handler_assign_test.go",
	}

	if mutate != nil {
		mutate(&capability)
	}

	return capability
}

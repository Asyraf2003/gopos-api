//go:build integration

package postgres

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

func TestCapabilityRepository_SaveGetAndList(t *testing.T) {
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	repository := NewCapabilityRepository(pool)
	txCtx := contextWithTx(ctx, tx)
	capability := integrationCapability(t, true)

	if err := repository.Save(txCtx, capability); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	found, err := repository.Get(txCtx, capability.Key)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if found.Key != capability.Key {
		t.Fatalf("key = %q, want %q", found.Key, capability.Key)
	}

	found = found.Disable("maintenance")
	if err := repository.Save(txCtx, found); err != nil {
		t.Fatalf("Save(disabled) error = %v", err)
	}

	disabled, err := repository.Get(txCtx, capability.Key)
	if err != nil {
		t.Fatalf("Get(disabled) error = %v", err)
	}
	if disabled.Enabled || disabled.DisabledReason != "maintenance" {
		t.Fatalf("disabled state = %v, reason = %q", disabled.Enabled, disabled.DisabledReason)
	}

	capabilities, err := repository.List(txCtx)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(capabilities) == 0 {
		t.Fatal("List() returned no capabilities")
	}
}

func TestCapabilityRepository_GetMissing(t *testing.T) {
	ctx := context.Background()
	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	repository := NewCapabilityRepository(pool)

	_, err := repository.Get(ctx, "missing.capability")
	if !errors.Is(err, ports.ErrCapabilityNotFound) {
		t.Fatalf("Get() error = %v, want not found", err)
	}
}

func integrationCapability(t *testing.T, enabled bool) domain.Capability {
	t.Helper()

	capability, err := domain.NewCapability(domain.Capability{
		Key:                "account.role.assign",
		Domain:             "account",
		Operation:          "assign-role",
		Method:             "POST",
		Path:               "/api/admin/accounts/:account_id/roles",
		DefaultEnabled:     true,
		Enabled:            enabled,
		RequiredPermission: "account.role.assign",
		RiskLevel:          domain.RiskHigh,
		AuditRequired:      true,
		OwnerPackage:       "internal/modules/auth",
		TestProof:          "internal/modules/auth/transport/http/account_role_handler_assign_test.go",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	return capability
}

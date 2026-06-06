package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"
)

type fakeCapabilityRepository struct {
	capabilities map[string]domain.Capability
	listCalls    int
	getCalls     int
	saveCalls    int
	lastSaved    domain.Capability
	err          error
}

func (f *fakeCapabilityRepository) List(ctx context.Context) ([]domain.Capability, error) {
	_ = ctx
	f.listCalls++
	if f.err != nil {
		return nil, f.err
	}

	out := make([]domain.Capability, 0, len(f.capabilities))
	for _, capability := range f.capabilities {
		out = append(out, capability)
	}

	return out, nil
}

func (f *fakeCapabilityRepository) Get(ctx context.Context, key string) (domain.Capability, error) {
	_ = ctx
	f.getCalls++
	if f.err != nil {
		return domain.Capability{}, f.err
	}

	capability, ok := f.capabilities[key]
	if !ok {
		return domain.Capability{}, ports.ErrCapabilityNotFound
	}

	return capability, nil
}

func (f *fakeCapabilityRepository) Save(ctx context.Context, capability domain.Capability) error {
	_ = ctx
	f.saveCalls++
	if f.err != nil {
		return f.err
	}

	f.lastSaved = capability
	f.capabilities[capability.Key] = capability

	return nil
}

func fakeRepository(t *testing.T, enabled bool) *fakeCapabilityRepository {
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
		DisabledReason:     "maintenance",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	return &fakeCapabilityRepository{
		capabilities: map[string]domain.Capability{
			capability.Key: capability,
		},
	}
}

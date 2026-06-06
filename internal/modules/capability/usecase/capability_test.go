package usecase

import (
	"context"
	"errors"
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

func TestCheckCapabilityAllowsEnabledCapability(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewCheckCapability(repository)

	err := usecase.Execute(context.Background(), "account.role.assign")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.getCalls != 1 {
		t.Fatalf("get calls = %d, want 1", repository.getCalls)
	}
}

func TestCheckCapabilityRejectsDisabledCapability(t *testing.T) {
	repository := fakeRepository(t, false)
	usecase := NewCheckCapability(repository)

	err := usecase.Execute(context.Background(), "account.role.assign")
	if !errors.Is(err, domain.ErrCapabilityDisabled) {
		t.Fatalf("Execute() error = %v, want disabled", err)
	}
}

func TestCheckCapabilityRejectsEmptyKeyBeforeRepository(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewCheckCapability(repository)

	err := usecase.Execute(context.Background(), " ")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
	if repository.getCalls != 0 {
		t.Fatalf("get calls = %d, want 0", repository.getCalls)
	}
}

func TestEnableCapabilityClearsDisabledReason(t *testing.T) {
	repository := fakeRepository(t, false)
	usecase := NewEnableCapability(repository)

	err := usecase.Execute(context.Background(), "account.role.assign")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !repository.lastSaved.Enabled {
		t.Fatal("saved capability is disabled")
	}
	if repository.lastSaved.DisabledReason != "" {
		t.Fatalf("disabled reason = %q, want empty", repository.lastSaved.DisabledReason)
	}
}

func TestDisableCapabilityStoresReason(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewDisableCapability(repository)

	err := usecase.Execute(context.Background(), "account.role.assign", "maintenance")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.lastSaved.Enabled {
		t.Fatal("saved capability is enabled")
	}
	if repository.lastSaved.DisabledReason != "maintenance" {
		t.Fatalf("disabled reason = %q", repository.lastSaved.DisabledReason)
	}
}

func TestShowCapabilityReturnsCapability(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewShowCapability(repository)

	capability, err := usecase.Execute(context.Background(), "account.role.assign")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if capability.Key != "account.role.assign" {
		t.Fatalf("key = %q", capability.Key)
	}
}

func TestListCapabilitiesReturnsAllCapabilities(t *testing.T) {
	repository := fakeRepository(t, true)
	usecase := NewListCapabilities(repository)

	capabilities, err := usecase.Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(capabilities) != 1 {
		t.Fatalf("capabilities len = %d, want 1", len(capabilities))
	}
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

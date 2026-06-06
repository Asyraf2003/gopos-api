package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/capability/domain"
)

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

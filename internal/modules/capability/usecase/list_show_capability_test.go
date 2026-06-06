package usecase

import (
	"context"
	"testing"
)

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

package memory

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func TestAuthStateStore_PutAndGetDel_SingleUse(t *testing.T) {
	store := NewAuthStateStore()

	err := store.Put(context.Background(), "state-123", ports.AuthState{
		Nonce:        "nonce-123",
		CodeVerifier: "verifier-123",
		Purpose:      "login",
		CreatedAt:    time.Now(),
	}, 5*time.Minute)
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	got, err := store.GetDel(context.Background(), "state-123")
	if err != nil {
		t.Fatalf("GetDel() first error = %v", err)
	}

	if got.Nonce != "nonce-123" {
		t.Fatalf("nonce = %q", got.Nonce)
	}
	if got.CodeVerifier != "verifier-123" {
		t.Fatalf("code verifier = %q", got.CodeVerifier)
	}
	if got.Purpose != "login" {
		t.Fatalf("purpose = %q", got.Purpose)
	}

	_, err = store.GetDel(context.Background(), "state-123")
	if err == nil {
		t.Fatal("GetDel() second call error = nil, want error")
	}
}

func TestAuthStateStore_GetDel_ExpiredState(t *testing.T) {
	store := NewAuthStateStore()

	err := store.Put(context.Background(), "state-expired", ports.AuthState{
		Nonce:        "nonce-expired",
		CodeVerifier: "verifier-expired",
		Purpose:      "login",
		CreatedAt:    time.Now(),
	}, 1*time.Nanosecond)
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	time.Sleep(1 * time.Millisecond)

	_, err = store.GetDel(context.Background(), "state-expired")
	if err == nil {
		t.Fatal("GetDel() expired error = nil, want error")
	}
}

func TestAuthStateStore_Put_RejectsInvalidInput(t *testing.T) {
	store := NewAuthStateStore()

	err := store.Put(context.Background(), "", ports.AuthState{}, 5*time.Minute)
	if err == nil {
		t.Fatal("Put() empty state error = nil, want error")
	}

	err = store.Put(context.Background(), "state-123", ports.AuthState{}, 0)
	if err == nil {
		t.Fatal("Put() zero ttl error = nil, want error")
	}
}

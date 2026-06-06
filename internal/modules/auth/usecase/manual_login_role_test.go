package usecase

import (
	"context"
	"testing"
	"time"
)

func TestManualLogin_KasirAssignsCashierRole(t *testing.T) {
	roles := &fakeManualRoleAssigner{}
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc-kasir"},
		roles,
		&fakeSessionStore{},
		&fakeTokenIssuer{token: "access-token", exp: time.Now().Add(15 * time.Minute)},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "kasir@example.com",
		Password: "12345678",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if roles.lastRoleKey != "cashier" {
		t.Fatalf("role key = %q, want cashier", roles.lastRoleKey)
	}
}

func TestManualLogin_RejectsUnsupportedEmail(t *testing.T) {
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc"},
		&fakeManualRoleAssigner{},
		&fakeSessionStore{},
		&fakeTokenIssuer{},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "owner@example.com",
		Password: "12345678",
	})
	if err != ErrManualLoginInvalidCredentials {
		t.Fatalf("error = %v, want ErrManualLoginInvalidCredentials", err)
	}
}

func TestManualLogin_RejectsInvalidPassword(t *testing.T) {
	usecase := NewManualLogin(
		&fakeManualAccountRepository{accountID: "acc"},
		&fakeManualRoleAssigner{},
		&fakeSessionStore{},
		&fakeTokenIssuer{},
		&fakeTransactor{},
		30*24*time.Hour,
	)

	_, err := usecase.Execute(context.Background(), ManualLoginInput{
		Email:    "admin@example.com",
		Password: "wrong-password",
	})
	if err != ErrManualLoginInvalidCredentials {
		t.Fatalf("error = %v, want ErrManualLoginInvalidCredentials", err)
	}
}

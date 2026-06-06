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
		Email: "kasir@example.com",
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
		Email: "owner@example.com",
	})
	if err != ErrManualLoginUnsupportedEmail {
		t.Fatalf("error = %v, want ErrManualLoginUnsupportedEmail", err)
	}
}

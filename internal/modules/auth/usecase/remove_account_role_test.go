package usecase

import (
	"context"
	"testing"
)

type fakeAccountRoleRemover struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeAccountRoleRemover) RemoveRole(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

func TestRemoveAccountRole_Success(t *testing.T) {
	remover := &fakeAccountRoleRemover{}
	usecase := NewRemoveAccountRole(remover)

	err := usecase.Execute(context.Background(), "acc-123", "admin")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if remover.calls != 1 {
		t.Fatalf("remover calls = %d, want 1", remover.calls)
	}
	if remover.lastAccountID != "acc-123" {
		t.Fatalf("account id = %q", remover.lastAccountID)
	}
	if remover.lastRoleKey != "admin" {
		t.Fatalf("role key = %q", remover.lastRoleKey)
	}
}

func TestRemoveAccountRole_RejectsBaseRole(t *testing.T) {
	remover := &fakeAccountRoleRemover{}
	usecase := NewRemoveAccountRole(remover)

	err := usecase.Execute(context.Background(), "acc-123", "base")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
	if err != ErrBaseRoleRemovalNotAllowed {
		t.Fatalf("error = %v", err)
	}
	if remover.calls != 0 {
		t.Fatalf("remover calls = %d, want 0", remover.calls)
	}
}

func TestRemoveAccountRole_RejectsEmptyAccountID(t *testing.T) {
	remover := &fakeAccountRoleRemover{}
	usecase := NewRemoveAccountRole(remover)

	err := usecase.Execute(context.Background(), "", "admin")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
	if remover.calls != 0 {
		t.Fatalf("remover calls = %d, want 0", remover.calls)
	}
}

func TestRemoveAccountRole_RejectsEmptyRoleKey(t *testing.T) {
	remover := &fakeAccountRoleRemover{}
	usecase := NewRemoveAccountRole(remover)

	err := usecase.Execute(context.Background(), "acc-123", "")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}
	if remover.calls != 0 {
		t.Fatalf("remover calls = %d, want 0", remover.calls)
	}
}

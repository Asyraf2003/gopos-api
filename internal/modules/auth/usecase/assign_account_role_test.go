package usecase

import (
	"context"
	"testing"
)

type fakeAssignAccountRoleAssigner struct {
	lastAccountID string
	lastRoleKey   string
	calls         int
	err           error
}

func (f *fakeAssignAccountRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.calls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return f.err
}

func TestAssignAccountRole_Success(t *testing.T) {
	assigner := &fakeAssignAccountRoleAssigner{}
	usecase := NewAssignAccountRole(assigner)

	err := usecase.Execute(context.Background(), "acc-123", "admin")
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if assigner.calls != 1 {
		t.Fatalf("assigner calls = %d, want 1", assigner.calls)
	}
	if assigner.lastAccountID != "acc-123" {
		t.Fatalf("account id = %q", assigner.lastAccountID)
	}
	if assigner.lastRoleKey != "admin" {
		t.Fatalf("role key = %q", assigner.lastRoleKey)
	}
}

func TestAssignAccountRole_RejectsEmptyAccountID(t *testing.T) {
	assigner := &fakeAssignAccountRoleAssigner{}
	usecase := NewAssignAccountRole(assigner)

	err := usecase.Execute(context.Background(), "", "admin")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}

	if assigner.calls != 0 {
		t.Fatalf("assigner calls = %d, want 0", assigner.calls)
	}
}

func TestAssignAccountRole_RejectsEmptyRoleKey(t *testing.T) {
	assigner := &fakeAssignAccountRoleAssigner{}
	usecase := NewAssignAccountRole(assigner)

	err := usecase.Execute(context.Background(), "acc-123", "")
	if err == nil {
		t.Fatal("Execute() error = nil, want error")
	}

	if assigner.calls != 0 {
		t.Fatalf("assigner calls = %d, want 0", assigner.calls)
	}
}

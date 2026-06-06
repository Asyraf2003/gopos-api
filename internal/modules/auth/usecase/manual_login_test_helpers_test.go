package usecase

import "context"

type fakeManualAccountRepository struct {
	accountID string
	lastEmail string
}

func (f *fakeManualAccountRepository) ResolveOrCreateManualAccount(ctx context.Context, email string) (string, error) {
	_ = ctx
	f.lastEmail = email
	return f.accountID, nil
}

type fakeManualRoleAssigner struct {
	ensureCalls   int
	lastAccountID string
	lastRoleKey   string
}

func (f *fakeManualRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	_ = ctx
	f.ensureCalls++
	f.lastAccountID = accountID
	f.lastRoleKey = roleKey
	return nil
}

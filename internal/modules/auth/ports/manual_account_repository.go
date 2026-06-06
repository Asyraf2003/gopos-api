package ports

import "context"

type ManualAccountRepository interface {
	ResolveOrCreateManualAccount(ctx context.Context, email string) (accountID string, err error)
}

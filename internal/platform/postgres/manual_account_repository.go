package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
)

func (r *AccountIdentityRepository) ResolveOrCreateManualAccount(ctx context.Context, email string) (string, error) {
	accountID, err := r.findAccountIDByEmail(ctx, email)
	if err == nil {
		return accountID, nil
	}
	if err != pgx.ErrNoRows {
		return "", err
	}

	return r.createAccount(ctx, email)
}

var _ ports.ManualAccountRepository = (*AccountIdentityRepository)(nil)

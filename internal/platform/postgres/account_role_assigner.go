package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRoleAssigner struct {
	pool *pgxpool.Pool
}

func NewAccountRoleAssigner(pool *pgxpool.Pool) *AccountRoleAssigner {
	return &AccountRoleAssigner{pool: pool}
}

func (a *AccountRoleAssigner) EnsureRole(ctx context.Context, accountID string, roleKey string) error {
	sql := `
		INSERT INTO account_roles (account_id, role_id)
		SELECT $1, r.id
		FROM roles r
		WHERE r.key = $2
		ON CONFLICT (account_id, role_id) DO NOTHING
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, accountID, roleKey)
		return err
	}

	_, err := a.pool.Exec(ctx, sql, accountID, roleKey)
	return err
}

var _ ports.AccountRoleAssigner = (*AccountRoleAssigner)(nil)

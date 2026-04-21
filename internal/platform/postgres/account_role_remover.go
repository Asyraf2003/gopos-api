package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRoleRemover struct {
	pool *pgxpool.Pool
}

func NewAccountRoleRemover(pool *pgxpool.Pool) *AccountRoleRemover {
	return &AccountRoleRemover{pool: pool}
}

func (r *AccountRoleRemover) RemoveRole(ctx context.Context, accountID string, roleKey string) error {
	sql := `
		DELETE FROM account_roles ar
		USING roles r
		WHERE ar.role_id = r.id
		  AND ar.account_id = $1
		  AND r.key = $2
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, accountID, roleKey)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, accountID, roleKey)
	return err
}

var _ ports.AccountRoleRemover = (*AccountRoleRemover)(nil)

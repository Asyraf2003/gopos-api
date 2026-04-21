package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRevoker struct {
	pool *pgxpool.Pool
}

func NewSessionRevoker(pool *pgxpool.Pool) *SessionRevoker {
	return &SessionRevoker{pool: pool}
}

func (r *SessionRevoker) RevokeSession(ctx context.Context, sessionID string) error {
	sql := `
		UPDATE auth_sessions
		SET revoked_at = now()
		WHERE id = $1
		  AND revoked_at IS NULL
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, sessionID)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, sessionID)
	return err
}

var _ ports.SessionRevoker = (*SessionRevoker)(nil)

package postgres

import (
	"context"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionStatusChecker struct {
	pool *pgxpool.Pool
}

func NewSessionStatusChecker(pool *pgxpool.Pool) *SessionStatusChecker {
	return &SessionStatusChecker{pool: pool}
}

func (s *SessionStatusChecker) IsSessionActive(ctx context.Context, sessionID string) (bool, error) {
	sql := `
		SELECT EXISTS (
			SELECT 1
			FROM auth_sessions
			WHERE id = $1
			  AND revoked_at IS NULL
		)
	`

	var active bool

	if tx, ok := TxFromContext(ctx); ok {
		err := tx.QueryRow(ctx, sql, sessionID).Scan(&active)
		return active, err
	}

	err := s.pool.QueryRow(ctx, sql, sessionID).Scan(&active)
	return active, err
}

var _ ports.SessionStatusChecker = (*SessionStatusChecker)(nil)

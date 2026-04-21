package postgres

import (
	"context"
	"encoding/json"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionStore struct {
	pool *pgxpool.Pool
}

func NewSessionStore(pool *pgxpool.Pool) *SessionStore {
	return &SessionStore{pool: pool}
}

func (s *SessionStore) Create(ctx context.Context, session domain.Session) error {
	metaJSON, err := json.Marshal(session.Meta)
	if err != nil {
		return err
	}

	sql := `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	args := []any{
		session.ID,
		session.AccountID,
		session.RefreshTokenHash,
		session.ExpiresAt,
		session.RevokedAt,
		metaJSON,
		session.CreatedAt,
	}

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, args...)
		return err
	}

	_, err = s.pool.Exec(ctx, sql, args...)
	return err
}

var _ ports.SessionStore = (*SessionStore)(nil)

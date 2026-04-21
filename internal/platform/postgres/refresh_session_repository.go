package postgres

import (
	"context"
	"time"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshSessionRepository struct {
	pool *pgxpool.Pool
}

func NewRefreshSessionRepository(pool *pgxpool.Pool) *RefreshSessionRepository {
	return &RefreshSessionRepository{pool: pool}
}

func (r *RefreshSessionRepository) FindActiveByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (ports.RefreshSession, error) {
	sql := `
		SELECT id, account_id, refresh_token_hash, expires_at, revoked_at
		FROM auth_sessions
		WHERE refresh_token_hash = $1
		  AND revoked_at IS NULL
		LIMIT 1
	`

	var session ports.RefreshSession

	if tx, ok := TxFromContext(ctx); ok {
		err := tx.QueryRow(ctx, sql, refreshTokenHash).Scan(
			&session.SessionID,
			&session.AccountID,
			&session.RefreshTokenHash,
			&session.ExpiresAt,
			&session.RevokedAt,
		)
		return session, err
	}

	err := r.pool.QueryRow(ctx, sql, refreshTokenHash).Scan(
		&session.SessionID,
		&session.AccountID,
		&session.RefreshTokenHash,
		&session.ExpiresAt,
		&session.RevokedAt,
	)
	return session, err
}

func (r *RefreshSessionRepository) RotateRefreshToken(ctx context.Context, sessionID string, newRefreshTokenHash string, newExpiresAt time.Time) error {
	sql := `
		UPDATE auth_sessions
		SET refresh_token_hash = $2,
		    expires_at = $3
		WHERE id = $1
		  AND revoked_at IS NULL
	`

	if tx, ok := TxFromContext(ctx); ok {
		_, err := tx.Exec(ctx, sql, sessionID, newRefreshTokenHash, newExpiresAt)
		return err
	}

	_, err := r.pool.Exec(ctx, sql, sessionID, newRefreshTokenHash, newExpiresAt)
	return err
}

var _ ports.RefreshSessionRepository = (*RefreshSessionRepository)(nil)

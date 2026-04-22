//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func mustInsertRefreshSessionFixture(
	t *testing.T,
	ctx context.Context,
	tx pgx.Tx,
) (string, string, string, time.Time) {
	t.Helper()

	accountID := uuid.NewString()
	sessionID := uuid.NewString()
	oldHash := "refresh-hash-old"
	oldExp := time.Now().Add(24 * time.Hour)

	_, err := tx.Exec(ctx, `
		INSERT INTO accounts (id, email, created_at, updated_at)
		VALUES ($1, $2, now(), now())
	`, accountID, "refresh-repo@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, NULL, '{}'::jsonb, now())
	`, sessionID, accountID, oldHash, oldExp)
	if err != nil {
		t.Fatalf("insert session error = %v", err)
	}

	return accountID, sessionID, oldHash, oldExp
}

func assertRefreshSessionFound(
	t *testing.T,
	got ports.RefreshSession,
	wantSessionID string,
	wantAccountID string,
	wantHash string,
) {
	t.Helper()

	if got.SessionID != wantSessionID {
		t.Fatalf("session id = %q, want %q", got.SessionID, wantSessionID)
	}
	if got.AccountID != wantAccountID {
		t.Fatalf("account id = %q, want %q", got.AccountID, wantAccountID)
	}
	if got.RefreshTokenHash != wantHash {
		t.Fatalf("refresh token hash = %q, want %q", got.RefreshTokenHash, wantHash)
	}
}

func mustQueryRefreshSessionRow(
	t *testing.T,
	ctx context.Context,
	tx pgx.Tx,
	sessionID string,
) ports.RefreshSession {
	t.Helper()

	var row ports.RefreshSession
	err := tx.QueryRow(ctx, `
		SELECT id, account_id, refresh_token_hash, expires_at, revoked_at
		FROM auth_sessions
		WHERE id = $1
	`, sessionID).Scan(
		&row.SessionID,
		&row.AccountID,
		&row.RefreshTokenHash,
		&row.ExpiresAt,
		&row.RevokedAt,
	)
	if err != nil {
		t.Fatalf("query rotated session error = %v", err)
	}

	return row
}

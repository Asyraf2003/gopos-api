//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestSessionRevoker_RevokeSession(t *testing.T) {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set")
	}

	ctx := context.Background()

	pool, err := NewPool(ctx, dsn)
	if err != nil {
		t.Fatalf("NewPool() error = %v", err)
	}
	defer pool.Close()

	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin() error = %v", err)
	}
	defer tx.Rollback(ctx)

	txCtx := context.WithValue(ctx, txContextKey{}, tx)

	accountID := uuid.NewString()
	_, err = tx.Exec(ctx, `
		INSERT INTO accounts (id, email, created_at, updated_at)
		VALUES ($1, $2, now(), now())
	`, accountID, "session-revoker@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	sessionID := uuid.NewString()
	_, err = tx.Exec(ctx, `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, NULL, '{}'::jsonb, now())
	`, sessionID, accountID, "hash-123", time.Now().Add(24*time.Hour))
	if err != nil {
		t.Fatalf("insert session error = %v", err)
	}

	revoker := NewSessionRevoker(pool)

	if err := revoker.RevokeSession(txCtx, sessionID); err != nil {
		t.Fatalf("RevokeSession() first call error = %v", err)
	}

	if err := revoker.RevokeSession(txCtx, sessionID); err != nil {
		t.Fatalf("RevokeSession() second call error = %v", err)
	}

	var revokedAt *time.Time
	err = tx.QueryRow(ctx, `
		SELECT revoked_at
		FROM auth_sessions
		WHERE id = $1
	`, sessionID).Scan(&revokedAt)
	if err != nil {
		t.Fatalf("query revoked_at error = %v", err)
	}

	if revokedAt == nil {
		t.Fatal("revoked_at is nil, want non-nil")
	}
}

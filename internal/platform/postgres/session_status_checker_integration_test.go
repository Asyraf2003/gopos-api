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

func TestSessionStatusChecker_IsSessionActive(t *testing.T) {
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
	`, accountID, "session-status-checker@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	activeSessionID := uuid.NewString()
	_, err = tx.Exec(ctx, `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, NULL, '{}'::jsonb, now())
	`, activeSessionID, accountID, "hash-active-123", time.Now().Add(24*time.Hour))
	if err != nil {
		t.Fatalf("insert active session error = %v", err)
	}

	revokedSessionID := uuid.NewString()
	_, err = tx.Exec(ctx, `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, now(), '{}'::jsonb, now())
	`, revokedSessionID, accountID, "hash-revoked-123", time.Now().Add(24*time.Hour))
	if err != nil {
		t.Fatalf("insert revoked session error = %v", err)
	}

	checker := NewSessionStatusChecker(pool)

	active, err := checker.IsSessionActive(txCtx, activeSessionID)
	if err != nil {
		t.Fatalf("IsSessionActive(active) error = %v", err)
	}
	if !active {
		t.Fatal("active session reported as inactive")
	}

	revoked, err := checker.IsSessionActive(txCtx, revokedSessionID)
	if err != nil {
		t.Fatalf("IsSessionActive(revoked) error = %v", err)
	}
	if revoked {
		t.Fatal("revoked session reported as active")
	}

	missing, err := checker.IsSessionActive(txCtx, uuid.NewString())
	if err != nil {
		t.Fatalf("IsSessionActive(missing) error = %v", err)
	}
	if missing {
		t.Fatal("missing session reported as active")
	}
}

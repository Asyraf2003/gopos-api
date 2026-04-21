//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestRefreshSessionRepository_FindAndRotate(t *testing.T) {
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
	`, accountID, "refresh-repo@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	sessionID := uuid.NewString()
	oldHash := "refresh-hash-old"
	oldExp := time.Now().Add(24 * time.Hour)

	_, err = tx.Exec(ctx, `
		INSERT INTO auth_sessions (
			id, account_id, refresh_token_hash, expires_at, revoked_at, meta_json, created_at
		) VALUES ($1, $2, $3, $4, NULL, '{}'::jsonb, now())
	`, sessionID, accountID, oldHash, oldExp)
	if err != nil {
		t.Fatalf("insert session error = %v", err)
	}

	repo := NewRefreshSessionRepository(pool)

	found, err := repo.FindActiveByRefreshTokenHash(txCtx, oldHash)
	if err != nil {
		t.Fatalf("FindActiveByRefreshTokenHash() error = %v", err)
	}

	if found.SessionID != sessionID {
		t.Fatalf("session id = %q, want %q", found.SessionID, sessionID)
	}
	if found.AccountID != accountID {
		t.Fatalf("account id = %q, want %q", found.AccountID, accountID)
	}
	if found.RefreshTokenHash != oldHash {
		t.Fatalf("refresh token hash = %q, want %q", found.RefreshTokenHash, oldHash)
	}

	newHash := "refresh-hash-new"
	newExp := time.Now().Add(48 * time.Hour)

	if err := repo.RotateRefreshToken(txCtx, sessionID, newHash, newExp); err != nil {
		t.Fatalf("RotateRefreshToken() error = %v", err)
	}

	foundAfter, err := repo.FindActiveByRefreshTokenHash(txCtx, newHash)
	if err != nil {
		t.Fatalf("FindActiveByRefreshTokenHash(new) error = %v", err)
	}

	if foundAfter.SessionID != sessionID {
		t.Fatalf("rotated session id = %q, want %q", foundAfter.SessionID, sessionID)
	}
	if foundAfter.RefreshTokenHash != newHash {
		t.Fatalf("rotated refresh token hash = %q, want %q", foundAfter.RefreshTokenHash, newHash)
	}
	if !foundAfter.ExpiresAt.After(found.ExpiresAt) {
		t.Fatalf("rotated expires_at = %v, want after %v", foundAfter.ExpiresAt, found.ExpiresAt)
	}

	_, err = repo.FindActiveByRefreshTokenHash(txCtx, oldHash)
	if err == nil {
		t.Fatal("old refresh token hash should not remain active")
	}

	var row ports.RefreshSession
	err = tx.QueryRow(ctx, `
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

	if row.RefreshTokenHash != newHash {
		t.Fatalf("db refresh token hash = %q, want %q", row.RefreshTokenHash, newHash)
	}
}

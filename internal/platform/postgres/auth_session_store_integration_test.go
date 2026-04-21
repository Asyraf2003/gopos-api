//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"pos-go/internal/modules/auth/domain"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestSessionStore_Create(t *testing.T) {
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
	`, accountID, "session-integration@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	store := NewSessionStore(pool)

	session := domain.Session{
		ID:               uuid.NewString(),
		AccountID:        accountID,
		RefreshTokenHash: "hash-1234567890abcdef",
		CreatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(24 * time.Hour),
		Meta: map[string]any{
			"provider": "google",
		},
	}

	if err := store.Create(txCtx, session); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	var (
		gotID               string
		gotAccountID        string
		gotRefreshTokenHash string
		gotRevokedAt        *time.Time
	)

	err = tx.QueryRow(ctx, `
		SELECT id, account_id, refresh_token_hash, revoked_at
		FROM auth_sessions
		WHERE id = $1
	`, session.ID).Scan(&gotID, &gotAccountID, &gotRefreshTokenHash, &gotRevokedAt)
	if err != nil {
		t.Fatalf("query session error = %v", err)
	}

	if gotID != session.ID {
		t.Fatalf("id = %q, want %q", gotID, session.ID)
	}
	if gotAccountID != session.AccountID {
		t.Fatalf("account_id = %q, want %q", gotAccountID, session.AccountID)
	}
	if gotRefreshTokenHash != session.RefreshTokenHash {
		t.Fatalf("refresh_token_hash = %q, want %q", gotRefreshTokenHash, session.RefreshTokenHash)
	}
	if gotRevokedAt != nil {
		t.Fatalf("revoked_at = %v, want nil", gotRevokedAt)
	}
}

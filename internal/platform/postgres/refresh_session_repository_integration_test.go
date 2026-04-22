//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"
)

func TestRefreshSessionRepository_FindAndRotate(t *testing.T) {
	ctx := context.Background()

	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	txCtx := contextWithTx(ctx, tx)
	accountID, sessionID, oldHash, oldExp := mustInsertRefreshSessionFixture(t, ctx, tx)

	repo := NewRefreshSessionRepository(pool)

	found, err := repo.FindActiveByRefreshTokenHash(txCtx, oldHash)
	if err != nil {
		t.Fatalf("FindActiveByRefreshTokenHash() error = %v", err)
	}
	assertRefreshSessionFound(t, found, sessionID, accountID, oldHash)

	newHash := "refresh-hash-new"
	newExp := time.Now().Add(48 * time.Hour)

	if err := repo.RotateRefreshToken(txCtx, sessionID, newHash, newExp); err != nil {
		t.Fatalf("RotateRefreshToken() error = %v", err)
	}

	foundAfter, err := repo.FindActiveByRefreshTokenHash(txCtx, newHash)
	if err != nil {
		t.Fatalf("FindActiveByRefreshTokenHash(new) error = %v", err)
	}
	assertRefreshSessionFound(t, foundAfter, sessionID, accountID, newHash)
	if !foundAfter.ExpiresAt.After(found.ExpiresAt) {
		t.Fatalf("rotated expires_at = %v, want after %v", foundAfter.ExpiresAt, found.ExpiresAt)
	}

	_, err = repo.FindActiveByRefreshTokenHash(txCtx, oldHash)
	if err == nil {
		t.Fatal("old refresh token hash should not remain active")
	}

	row := mustQueryRefreshSessionRow(t, ctx, tx, sessionID)
	if row.RefreshTokenHash != newHash {
		t.Fatalf("db refresh token hash = %q, want %q", row.RefreshTokenHash, newHash)
	}

	_ = oldExp
}

//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/domain"
)

func TestAccountIdentityRepository_ResolveCreateAndUpsert(t *testing.T) {
	ctx := context.Background()

	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	txCtx := contextWithTx(ctx, tx)
	repo := NewAccountIdentityRepository(pool)
	claims := newAccountIdentityTestClaims()

	accountID, err := repo.ResolveOrCreateAccountByGoogle(txCtx, claims)
	if err != nil {
		t.Fatalf("ResolveOrCreateAccountByGoogle() error = %v", err)
	}
	if accountID == "" {
		t.Fatal("accountID is empty")
	}

	assertAccountEmailByID(t, ctx, tx, accountID, claims.Email)

	accountID2, err := repo.ResolveOrCreateAccountByGoogle(txCtx, claims)
	if err != nil {
		t.Fatalf("ResolveOrCreateAccountByGoogle() second call error = %v", err)
	}
	if accountID2 != accountID {
		t.Fatalf("second account id = %q, want %q", accountID2, accountID)
	}

	identity := domain.Identity{
		Provider:      domain.ProviderGoogle,
		Subject:       claims.Subject,
		Email:         claims.Email,
		EmailVerified: true,
		Meta: map[string]any{
			"source": "integration-test",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.UpsertIdentity(txCtx, accountID, identity); err != nil {
		t.Fatalf("UpsertIdentity() error = %v", err)
	}

	assertIdentityRow(t, ctx, tx, accountID, claims)
}

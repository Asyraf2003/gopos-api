//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/joho/godotenv"
)

func TestAccountIdentityRepository_ResolveCreateAndUpsert(t *testing.T) {
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

	repo := NewAccountIdentityRepository(pool)

	claims := ports.OIDCClaims{
		Provider:      "google",
		Subject:       "integration-subject-123",
		Email:         "integration-user@example.com",
		EmailVerified: true,
		AuthTime:      time.Now(),
	}

	accountID, err := repo.ResolveOrCreateAccountByGoogle(txCtx, claims)
	if err != nil {
		t.Fatalf("ResolveOrCreateAccountByGoogle() error = %v", err)
	}
	if accountID == "" {
		t.Fatal("accountID is empty")
	}

	var accountEmail string
	err = tx.QueryRow(ctx, `SELECT email FROM accounts WHERE id = $1`, accountID).Scan(&accountEmail)
	if err != nil {
		t.Fatalf("query account error = %v", err)
	}
	if accountEmail != claims.Email {
		t.Fatalf("account email = %q, want %q", accountEmail, claims.Email)
	}

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

	var gotProvider, gotSubject, gotEmail string
	var gotVerified bool
	err = tx.QueryRow(ctx, `
		SELECT provider, subject, email, email_verified
		FROM auth_identities
		WHERE account_id = $1 AND provider = $2 AND subject = $3
	`,
		accountID,
		"google",
		claims.Subject,
	).Scan(&gotProvider, &gotSubject, &gotEmail, &gotVerified)
	if err != nil {
		t.Fatalf("query identity error = %v", err)
	}

	if gotProvider != "google" {
		t.Fatalf("provider = %q", gotProvider)
	}
	if gotSubject != claims.Subject {
		t.Fatalf("subject = %q, want %q", gotSubject, claims.Subject)
	}
	if gotEmail != claims.Email {
		t.Fatalf("email = %q, want %q", gotEmail, claims.Email)
	}
	if !gotVerified {
		t.Fatal("email_verified = false, want true")
	}
}

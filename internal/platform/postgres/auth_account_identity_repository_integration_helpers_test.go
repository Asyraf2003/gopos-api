//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/auth/ports"

	"github.com/jackc/pgx/v5"
)

func newAccountIdentityTestClaims() ports.OIDCClaims {
	return ports.OIDCClaims{
		Provider:      "google",
		Subject:       "integration-subject-123",
		Email:         "integration-user@example.com",
		EmailVerified: true,
		AuthTime:      time.Now(),
	}
}

func assertAccountEmailByID(t *testing.T, ctx context.Context, tx pgx.Tx, accountID string, wantEmail string) {
	t.Helper()

	var gotEmail string
	err := tx.QueryRow(ctx, `SELECT email FROM accounts WHERE id = $1`, accountID).Scan(&gotEmail)
	if err != nil {
		t.Fatalf("query account error = %v", err)
	}
	if gotEmail != wantEmail {
		t.Fatalf("account email = %q, want %q", gotEmail, wantEmail)
	}
}

func assertIdentityRow(t *testing.T, ctx context.Context, tx pgx.Tx, accountID string, claims ports.OIDCClaims) {
	t.Helper()

	var gotProvider string
	var gotSubject string
	var gotEmail string
	var gotVerified bool

	err := tx.QueryRow(ctx, `
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

//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func mustOpenIntegrationPool(t *testing.T, ctx context.Context) *pgxpool.Pool {
	t.Helper()

	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set")
	}

	pool, err := NewPool(ctx, dsn)
	if err != nil {
		t.Fatalf("NewPool() error = %v", err)
	}

	return pool
}

func mustBeginIntegrationTx(t *testing.T, ctx context.Context, pool *pgxpool.Pool) pgx.Tx {
	t.Helper()

	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin() error = %v", err)
	}

	return tx
}

func contextWithTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txContextKey{}, tx)
}

func mustInsertPrincipalResolverFixture(t *testing.T, ctx context.Context, tx pgx.Tx) string {
	t.Helper()

	accountID := uuid.NewString()

	_, err := tx.Exec(ctx, `
		INSERT INTO accounts (id, email, created_at, updated_at)
		VALUES ($1, $2, now(), now())
	`, accountID, "principal-resolver@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO account_roles (account_id, role_id)
		SELECT $1, r.id
		FROM roles r
		WHERE r.key IN ('base', 'cashier')
		ON CONFLICT (account_id, role_id) DO NOTHING
	`, accountID)
	if err != nil {
		t.Fatalf("insert account_roles error = %v", err)
	}

	return accountID
}

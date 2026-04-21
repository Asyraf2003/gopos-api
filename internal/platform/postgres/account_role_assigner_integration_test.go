//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestAccountRoleAssigner_EnsureRole(t *testing.T) {
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
	`, accountID, "role-assigner-integration@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	assigner := NewAccountRoleAssigner(pool)

	if err := assigner.EnsureRole(txCtx, accountID, "base"); err != nil {
		t.Fatalf("EnsureRole() first call error = %v", err)
	}

	if err := assigner.EnsureRole(txCtx, accountID, "base"); err != nil {
		t.Fatalf("EnsureRole() second call error = %v", err)
	}

	var count int
	err = tx.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1 AND r.key = 'base'
	`, accountID).Scan(&count)
	if err != nil {
		t.Fatalf("count query error = %v", err)
	}

	if count != 1 {
		t.Fatalf("base role count = %d, want 1", count)
	}
}

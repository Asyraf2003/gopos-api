//go:build integration

package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestAccountRoleRemover_RemoveRole(t *testing.T) {
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
	roleKey := "role-remove-test-" + uuid.NewString()

	_, err = tx.Exec(ctx, `
		INSERT INTO accounts (id, email, created_at, updated_at)
		VALUES ($1, $2, now(), now())
	`, accountID, "account-role-remover@example.com")
	if err != nil {
		t.Fatalf("insert account error = %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO roles (key, name, created_at)
		VALUES ($1, $2, now())
	`, roleKey, "Role Remove Test")
	if err != nil {
		t.Fatalf("insert role error = %v", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO account_roles (account_id, role_id, created_at)
		SELECT $1, r.id, now()
		FROM roles r
		WHERE r.key = $2
	`, accountID, roleKey)
	if err != nil {
		t.Fatalf("insert account role error = %v", err)
	}

	var countBefore int
	err = tx.QueryRow(ctx, `
		SELECT count(*)
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		  AND r.key = $2
	`, accountID, roleKey).Scan(&countBefore)
	if err != nil {
		t.Fatalf("count before error = %v", err)
	}
	if countBefore != 1 {
		t.Fatalf("count before = %d, want 1", countBefore)
	}

	remover := NewAccountRoleRemover(pool)

	if err := remover.RemoveRole(txCtx, accountID, roleKey); err != nil {
		t.Fatalf("RemoveRole() first call error = %v", err)
	}

	if err := remover.RemoveRole(txCtx, accountID, roleKey); err != nil {
		t.Fatalf("RemoveRole() second call error = %v", err)
	}

	var countAfter int
	err = tx.QueryRow(ctx, `
		SELECT count(*)
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		  AND r.key = $2
	`, accountID, roleKey).Scan(&countAfter)
	if err != nil {
		t.Fatalf("count after error = %v", err)
	}
	if countAfter != 0 {
		t.Fatalf("count after = %d, want 0", countAfter)
	}
}

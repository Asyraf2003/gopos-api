//go:build integration

package postgres

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func mustInsertAccountRoleRemoverFixture(t *testing.T, ctx context.Context, tx pgx.Tx) (string, string) {
	t.Helper()

	accountID := uuid.NewString()
	roleKey := "role-remove-test-" + uuid.NewString()

	_, err := tx.Exec(ctx, `
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

	return accountID, roleKey
}

func mustCountAccountRoleByKey(
	t *testing.T,
	ctx context.Context,
	tx pgx.Tx,
	accountID string,
	roleKey string,
) int {
	t.Helper()

	var count int
	err := tx.QueryRow(ctx, `
		SELECT count(*)
		FROM account_roles ar
		JOIN roles r ON r.id = ar.role_id
		WHERE ar.account_id = $1
		  AND r.key = $2
	`, accountID, roleKey).Scan(&count)
	if err != nil {
		t.Fatalf("count account role error = %v", err)
	}

	return count
}

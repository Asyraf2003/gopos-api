//go:build integration

package postgres

import (
	"context"
	"os"
	"reflect"
	"testing"

	"pos-go/internal/modules/auth/ports"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestPrincipalResolver_Resolve(t *testing.T) {
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

	resolver := NewPrincipalResolver(pool)

	principal, err := resolver.Resolve(txCtx, ports.ResolvePrincipalInput{
		AccountID:  accountID,
		SessionID:  "session-123",
		TrustLevel: "aal1",
	})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if principal.AccountID != accountID {
		t.Fatalf("account_id = %q, want %q", principal.AccountID, accountID)
	}
	if principal.SessionID != "session-123" {
		t.Fatalf("session_id = %q, want session-123", principal.SessionID)
	}
	if principal.TrustLevel != "aal1" {
		t.Fatalf("trust_level = %q, want aal1", principal.TrustLevel)
	}

	wantRoles := []string{"base", "cashier"}
	if !reflect.DeepEqual(principal.Roles, wantRoles) {
		t.Fatalf("roles = %#v, want %#v", principal.Roles, wantRoles)
	}

	wantPermissions := []string{
		"auth.session.logout",
		"auth.session.refresh",
		"payment.create",
		"profile.self.read",
		"sale.order.create",
		"sale.order.read",
	}
	if !reflect.DeepEqual(principal.Permissions, wantPermissions) {
		t.Fatalf("permissions = %#v, want %#v", principal.Permissions, wantPermissions)
	}

	if !principal.HasPermission("sale.order.create") {
		t.Fatal("expected permission sale.order.create")
	}
	if principal.HasPermission("inventory.manage") {
		t.Fatal("did not expect permission inventory.manage")
	}
}

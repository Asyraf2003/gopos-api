//go:build integration

package postgres

import (
	"context"
	"testing"
)

func TestAccountRoleRemover_RemoveRole(t *testing.T) {
	ctx := context.Background()

	pool := mustOpenIntegrationPool(t, ctx)
	defer pool.Close()

	tx := mustBeginIntegrationTx(t, ctx, pool)
	defer tx.Rollback(ctx)

	txCtx := contextWithTx(ctx, tx)
	accountID, roleKey := mustInsertAccountRoleRemoverFixture(t, ctx, tx)

	countBefore := mustCountAccountRoleByKey(t, ctx, tx, accountID, roleKey)
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

	countAfter := mustCountAccountRoleByKey(t, ctx, tx, accountID, roleKey)
	if countAfter != 0 {
		t.Fatalf("count after = %d, want 0", countAfter)
	}
}

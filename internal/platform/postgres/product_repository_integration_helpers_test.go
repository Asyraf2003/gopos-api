//go:build integration

package postgres

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func mustOpenProductCatalogRepoTx(
	t *testing.T,
	ctx context.Context,
) (*pgxpool.Pool, context.Context) {
	t.Helper()

	pool := mustOpenIntegrationPool(t, ctx)
	tx := mustBeginIntegrationTx(t, ctx, pool)
	t.Cleanup(func() {
		_ = tx.Rollback(ctx)
		pool.Close()
	})

	return pool, contextWithTx(ctx, tx)
}

func newProductCatalogTestProduct(t *testing.T, name string) *domain.Product {
	t.Helper()

	product, err := domain.NewProduct(domain.ProductInput{
		ID:              uuid.NewString(),
		Code:            "SKU-" + uuid.NewString(),
		Name:            name,
		Brand:           "Honda",
		SalePriceRupiah: 40000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	return product
}

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductReturnsDuplicateCheckerError(t *testing.T) {
	duplicateErr := errors.New("duplicate update failure")
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_002",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{err: duplicateErr},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_002",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})

	if !errors.Is(err, duplicateErr) {
		t.Fatalf("Execute() error = %v, want %v", err, duplicateErr)
	}
}

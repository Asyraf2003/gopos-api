package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestUpdateProductReturnsNotFoundWhenProductDoesNotExist(t *testing.T) {
	uc := NewUpdateProduct(
		&fakeProductRepository{},
		&fakeProductDuplicateChecker{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		time.Now,
	)

	_, err := uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "missing_product",
		Name:            "Busi",
		Brand:           "NGK",
		SalePriceRupiah: 25000,
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("Execute() error = %v, want %v", err, ports.ErrProductNotFound)
	}
}

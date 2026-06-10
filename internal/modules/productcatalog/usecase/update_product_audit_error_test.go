package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestUpdateProductReturnsAuditRecorderError(t *testing.T) {
	auditErr := errors.New("audit record failure")
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_008",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	auditRecorder := &fakeProductAuditRecorder{err: auditErr}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{},
		&fakeUpdateProductVersionRepository{},
		auditRecorder,
		time.Now,
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_008",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
	})

	if !errors.Is(err, auditErr) {
		t.Fatalf("Execute() error = %v, want %v", err, auditErr)
	}
	if len(auditRecorder.records) != 1 {
		t.Fatalf("audit record count = %d, want 1", len(auditRecorder.records))
	}
}

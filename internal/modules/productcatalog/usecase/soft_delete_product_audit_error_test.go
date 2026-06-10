package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestSoftDeleteProductPropagatesAuditRecordError(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	auditErr := errors.New("audit record failed")
	usecase := NewSoftDeleteProduct(
		&softDeleteProductRepositoryDouble{found: product},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{err: auditErr},
		func() time.Time { return time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC) },
	)

	_, err = usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-1",
		ActorID: "actor-1",
		Reason:  "obsolete product",
	})

	if !errors.Is(err, auditErr) {
		t.Fatalf("expected audit record error, got %v", err)
	}
}

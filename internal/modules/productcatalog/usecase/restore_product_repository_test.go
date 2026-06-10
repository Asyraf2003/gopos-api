package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestRestoreProductUpdatesRepository(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 10, 9, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor-1",
		Reason:           "setup deleted product",
	}); err != nil {
		t.Fatalf("SoftDelete() setup error = %v", err)
	}

	repository := &restoreProductRepositoryDouble{found: product}
	usecase := NewRestoreProduct(
		repository,
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		func() time.Time { return time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC) },
	)

	_, err = usecase.Execute(context.Background(), RestoreProductCommand{
		ID:      "product-1",
		ActorID: "actor-2",
		Reason:  "restore product",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.updated == nil {
		t.Fatal("expected repository update, got nil")
	}
	if repository.updated.Status() != domain.ProductStatusActive {
		t.Fatalf("Status() = %v, want %v", repository.updated.Status(), domain.ProductStatusActive)
	}
	if repository.updated.DeletedAt() != nil {
		t.Fatalf("DeletedAt() = %v, want nil", repository.updated.DeletedAt())
	}
}

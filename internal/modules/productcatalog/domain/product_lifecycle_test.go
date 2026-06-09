package domain

import (
	"testing"
	"time"
)

func TestProductSoftDeleteAndRestoreLifecycle(t *testing.T) {
	product, err := NewProduct(ProductInput{
		ID:              "prod_003",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	deletedAt := time.Date(2026, 6, 9, 10, 0, 0, 0, time.UTC)
	if err := product.SoftDelete(DeleteInput{
		DeletedAt:        deletedAt,
		DeletedByActorID: "actor_admin",
		Reason:           "duplicate test product",
	}); err != nil {
		t.Fatalf("SoftDelete() error = %v", err)
	}

	if product.Status() != ProductStatusDeleted {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusDeleted)
	}
	if product.DeletedAt() == nil || !product.DeletedAt().Equal(deletedAt) {
		t.Fatalf("DeletedAt() = %v, want %v", product.DeletedAt(), deletedAt)
	}

	if err := product.Restore(); err != nil {
		t.Fatalf("Restore() error = %v", err)
	}

	if product.Status() != ProductStatusActive {
		t.Fatalf("Status() = %v, want %v", product.Status(), ProductStatusActive)
	}
	if product.DeletedAt() != nil {
		t.Fatalf("DeletedAt() = %v, want nil", product.DeletedAt())
	}
}

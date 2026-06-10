package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestRestoreProductPropagatesVersionAppendError(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	if err := product.SoftDelete(domain.DeleteInput{
		DeletedAt:        time.Date(2026, 6, 10, 9, 0, 0, 0, time.UTC),
		DeletedByActorID: "actor-1",
		Reason:           "setup deleted product",
	}); err != nil {
		t.Fatalf("SoftDelete() setup error = %v", err)
	}

	versionErr := errors.New("version append failed")
	usecase := NewRestoreProduct(
		&restoreProductRepositoryDouble{found: product},
		&fakeProductVersionRepository{err: versionErr},
		&fakeProductAuditRecorder{},
		func() time.Time { return time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC) },
	)

	_, err = usecase.Execute(context.Background(), RestoreProductCommand{
		ID:      "product-1",
		ActorID: "actor-2",
		Reason:  "restore product",
	})

	if !errors.Is(err, versionErr) {
		t.Fatalf("expected version append error, got %v", err)
	}
}

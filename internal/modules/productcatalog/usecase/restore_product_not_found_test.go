package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestRestoreProductReturnsNotFound(t *testing.T) {
	usecase := NewRestoreProduct(
		&fakeProductRepository{},
		&fakeProductVersionRepository{},
		&fakeProductAuditRecorder{},
		func() time.Time { return time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC) },
	)

	_, err := usecase.Execute(context.Background(), RestoreProductCommand{
		ID:      "product-404",
		ActorID: "actor-1",
		Reason:  "restore missing product",
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("expected product not found, got %v", err)
	}
}

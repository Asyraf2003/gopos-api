package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestSoftDeleteProductReturnsNotFound(t *testing.T) {
	repository := &productRepositoryDouble{
		findErr: ports.ErrProductNotFound,
	}
	usecase := NewSoftDeleteProduct(
		repository,
		&productVersionRepositoryDouble{},
		&productAuditRecorderDouble{},
		func() time.Time { return time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC) },
	)

	_, err := usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-404",
		ActorID: "actor-1",
		Reason:  "delete missing product",
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("expected product not found, got %v", err)
	}
}

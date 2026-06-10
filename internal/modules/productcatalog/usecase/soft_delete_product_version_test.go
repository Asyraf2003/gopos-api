package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
)

func TestSoftDeleteProductRecordsVersion(t *testing.T) {
	product, err := domain.NewProduct(domain.ProductInput{
		ID:              "product-1",
		Name:            "Filter Udara",
		Brand:           "Aspira",
		SalePriceRupiah: 30000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	versionRepository := &fakeProductVersionRepository{}
	deletedAt := time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC)
	usecase := NewSoftDeleteProduct(
		&softDeleteProductRepositoryDouble{found: product},
		versionRepository,
		&fakeProductAuditRecorder{},
		func() time.Time { return deletedAt },
	)

	_, err = usecase.Execute(context.Background(), SoftDeleteProductCommand{
		ID:      "product-1",
		ActorID: "actor-1",
		Reason:  "obsolete product",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(versionRepository.appended) != 1 {
		t.Fatalf("version append count = %d, want 1", len(versionRepository.appended))
	}

	version := versionRepository.appended[0]
	if version.ProductID != "product-1" {
		t.Fatalf("ProductID = %q, want product-1", version.ProductID)
	}
	if version.EventName != "product_deleted" {
		t.Fatalf("EventName = %q, want product_deleted", version.EventName)
	}
	if version.ChangedByActorID != "actor-1" {
		t.Fatalf("ChangedByActorID = %q, want actor-1", version.ChangedByActorID)
	}
	if version.ChangeReason != "obsolete product" {
		t.Fatalf("ChangeReason = %q, want obsolete product", version.ChangeReason)
	}
	if !version.ChangedAt.Equal(deletedAt) {
		t.Fatalf("ChangedAt = %v, want %v", version.ChangedAt, deletedAt)
	}
	if version.RevisionNo != 1 {
		t.Fatalf("RevisionNo = %d, want 1", version.RevisionNo)
	}
}

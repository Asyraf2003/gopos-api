package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestUpdateProductRecordsVersion(t *testing.T) {
	fixedNow := time.Date(2026, 6, 10, 10, 0, 0, 0, time.UTC)
	existing, err := domain.NewProduct(domain.ProductInput{
		ID:              "prod_005",
		Name:            "Busi Lama",
		Brand:           "NGK",
		SalePriceRupiah: 20000,
	})
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	versionRepository := &fakeUpdateProductVersionRepository{
		existing: []ports.ProductVersionRecord{{RevisionNo: 1}},
	}
	uc := NewUpdateProduct(
		&fakeUpdateProductRepository{product: existing},
		&fakeUpdateProductDuplicateChecker{},
		versionRepository,
		&fakeProductAuditRecorder{},
		func() time.Time { return fixedNow },
	)

	_, err = uc.Execute(context.Background(), UpdateProductCommand{
		ID:              "prod_005",
		Name:            "Busi Baru",
		Brand:           "Denso",
		SalePriceRupiah: 30000,
		ActorID:         "actor_001",
		Reason:          "price correction",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(versionRepository.appended) != 1 {
		t.Fatalf("version append count = %d, want 1", len(versionRepository.appended))
	}
	version := versionRepository.appended[0]
	if version.ProductID != "prod_005" {
		t.Fatalf("version.ProductID = %q, want prod_005", version.ProductID)
	}
	if version.RevisionNo != 2 {
		t.Fatalf("version.RevisionNo = %d, want 2", version.RevisionNo)
	}
	if version.EventName != productUpdatedEventName {
		t.Fatalf("version.EventName = %q, want %q", version.EventName, productUpdatedEventName)
	}
	if version.ChangedByActorID != "actor_001" {
		t.Fatalf("version.ChangedByActorID = %q, want actor_001", version.ChangedByActorID)
	}
	if version.ChangeReason != "price correction" {
		t.Fatalf("version.ChangeReason = %q, want price correction", version.ChangeReason)
	}
	if !version.ChangedAt.Equal(fixedNow) {
		t.Fatalf("version.ChangedAt = %v, want %v", version.ChangedAt, fixedNow)
	}
}

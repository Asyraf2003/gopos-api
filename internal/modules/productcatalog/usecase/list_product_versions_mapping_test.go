package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductVersionsMapsRecords(t *testing.T) {
	changedAt := time.Date(2026, 6, 10, 11, 0, 0, 0, time.UTC)
	usecase := NewListProductVersions(&listProductVersionsRepositoryDouble{
		records: []ports.ProductVersionRecord{
			{
				ProductID:        "prod_001",
				RevisionNo:       2,
				EventName:        "product.updated",
				ChangedByActorID: "actor_001",
				ChangeReason:     "price correction",
				ChangedAt:        changedAt,
			},
		},
	})

	result, err := usecase.Execute(context.Background(), ListProductVersionsQuery{
		ProductID: "prod_001",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 1 {
		t.Fatalf("Items length = %d, want 1", len(result.Items))
	}

	item := result.Items[0]
	if item.ProductID != "prod_001" {
		t.Fatalf("ProductID = %q, want prod_001", item.ProductID)
	}
	if item.RevisionNo != 2 {
		t.Fatalf("RevisionNo = %d, want 2", item.RevisionNo)
	}
	if item.EventName != "product.updated" {
		t.Fatalf("EventName = %q, want product.updated", item.EventName)
	}
	if item.ChangedByActorID != "actor_001" {
		t.Fatalf("ChangedByActorID = %q, want actor_001", item.ChangedByActorID)
	}
	if item.ChangeReason != "price correction" {
		t.Fatalf("ChangeReason = %q, want price correction", item.ChangeReason)
	}
	if !item.ChangedAt.Equal(changedAt) {
		t.Fatalf("ChangedAt = %v, want %v", item.ChangedAt, changedAt)
	}
}

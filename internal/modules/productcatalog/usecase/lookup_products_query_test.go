package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestLookupProductsForwardsQuery(t *testing.T) {
	reader := &lookupProductsReaderDouble{}
	usecase := NewLookupProducts(reader)

	_, err := usecase.Execute(context.Background(), LookupProductsQuery{
		Query:          "filter",
		Limit:          15,
		IncludeDeleted: true,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if reader.capturedQuery.Query != "filter" {
		t.Fatalf("Query = %q, want filter", reader.capturedQuery.Query)
	}
	if reader.capturedQuery.Limit != 15 {
		t.Fatalf("Limit = %d, want 15", reader.capturedQuery.Limit)
	}
	if !reader.capturedQuery.IncludeDeleted {
		t.Fatalf("IncludeDeleted = false, want true")
	}
}

var _ ports.ProductReader = (*lookupProductsReaderDouble)(nil)

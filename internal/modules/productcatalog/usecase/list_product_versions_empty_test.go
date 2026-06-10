package usecase

import (
	"context"
	"testing"
)

func TestListProductVersionsReturnsEmptyItems(t *testing.T) {
	usecase := NewListProductVersions(&listProductVersionsRepositoryDouble{})

	result, err := usecase.Execute(context.Background(), ListProductVersionsQuery{
		ProductID: "prod_001",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(result.Items) != 0 {
		t.Fatalf("Items length = %d, want 0", len(result.Items))
	}
	if result.Items == nil {
		t.Fatalf("Items = nil, want empty slice")
	}
}

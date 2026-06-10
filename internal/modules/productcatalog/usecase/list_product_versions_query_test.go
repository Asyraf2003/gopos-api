package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductVersionsForwardsProductID(t *testing.T) {
	repository := &listProductVersionsRepositoryDouble{}
	usecase := NewListProductVersions(repository)

	_, err := usecase.Execute(context.Background(), ListProductVersionsQuery{
		ProductID: "prod_001",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if repository.capturedProductID != "prod_001" {
		t.Fatalf("ProductID = %q, want prod_001", repository.capturedProductID)
	}
}

var _ ports.ProductVersionRepository = (*listProductVersionsRepositoryDouble)(nil)

package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductVersionsPropagatesRepositoryError(t *testing.T) {
	listErr := errors.New("list product versions failed")
	usecase := NewListProductVersions(&listProductVersionsRepositoryDouble{
		listErr: listErr,
	})

	_, err := usecase.Execute(context.Background(), ListProductVersionsQuery{
		ProductID: "prod_001",
	})

	if !errors.Is(err, listErr) {
		t.Fatalf("Execute() error = %v, want %v", err, listErr)
	}
}

type listProductVersionsRepositoryDouble struct {
	capturedProductID string
	records           []ports.ProductVersionRecord
	listErr           error
	appended          []ports.ProductVersionRecord
}

func (d *listProductVersionsRepositoryDouble) Append(
	_ context.Context,
	version ports.ProductVersionRecord,
) error {
	d.appended = append(d.appended, version)

	return nil
}

func (d *listProductVersionsRepositoryDouble) ListByProductID(
	_ context.Context,
	productID string,
) ([]ports.ProductVersionRecord, error) {
	d.capturedProductID = productID
	if d.listErr != nil {
		return nil, d.listErr
	}

	return d.records, nil
}

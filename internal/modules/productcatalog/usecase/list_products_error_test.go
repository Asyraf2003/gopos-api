package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestListProductsPropagatesReaderListError(t *testing.T) {
	listErr := errors.New("list products failed")
	usecase := NewListProducts(&listProductsReaderDouble{
		listErr: listErr,
	})

	_, err := usecase.Execute(context.Background(), ListProductsQuery{})

	if !errors.Is(err, listErr) {
		t.Fatalf("Execute() error = %v, want %v", err, listErr)
	}
}

type listProductsReaderDouble struct {
	listErr error
}

func (d *listProductsReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return nil, nil
}

func (d *listProductsReaderDouble) List(
	_ context.Context,
	_ ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	return nil, d.listErr
}

func (d *listProductsReaderDouble) Lookup(
	_ context.Context,
	_ ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	return nil, nil
}

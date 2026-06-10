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
	listErr       error
	listItems     []ports.ProductListItem
	capturedQuery ports.ProductListQuery
}

func (d *listProductsReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return nil, nil
}

func (d *listProductsReaderDouble) List(
	_ context.Context,
	query ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	d.capturedQuery = query

	return d.listItems, d.listErr
}

func (d *listProductsReaderDouble) Lookup(
	_ context.Context,
	_ ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	return nil, nil
}

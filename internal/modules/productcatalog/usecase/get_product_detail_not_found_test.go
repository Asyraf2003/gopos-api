package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestGetProductDetailReturnsNotFound(t *testing.T) {
	usecase := NewGetProductDetail(&getProductDetailReaderDouble{
		err: ports.ErrProductNotFound,
	})

	_, err := usecase.Execute(context.Background(), GetProductDetailQuery{
		ID: "product-404",
	})

	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("expected product not found, got %v", err)
	}
}

type getProductDetailReaderDouble struct {
	found *domain.Product
	err   error
}

func (d *getProductDetailReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	if d.err != nil {
		return nil, d.err
	}

	return d.found, nil
}

func (d *getProductDetailReaderDouble) List(
	_ context.Context,
	_ ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	return nil, nil
}

func (d *getProductDetailReaderDouble) Lookup(
	_ context.Context,
	_ ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	return nil, nil
}

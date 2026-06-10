package usecase

import (
	"context"
	"errors"
	"testing"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

func TestLookupProductsPropagatesReaderLookupError(t *testing.T) {
	lookupErr := errors.New("lookup products failed")
	usecase := NewLookupProducts(&lookupProductsReaderDouble{
		lookupErr: lookupErr,
	})

	_, err := usecase.Execute(context.Background(), LookupProductsQuery{})

	if !errors.Is(err, lookupErr) {
		t.Fatalf("Execute() error = %v, want %v", err, lookupErr)
	}
}

type lookupProductsReaderDouble struct {
	capturedQuery ports.ProductLookupQuery
	lookupErr     error
}

func (d *lookupProductsReaderDouble) GetByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return nil, nil
}

func (d *lookupProductsReaderDouble) List(
	_ context.Context,
	_ ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	return nil, nil
}

func (d *lookupProductsReaderDouble) Lookup(
	_ context.Context,
	query ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	d.capturedQuery = query

	return nil, d.lookupErr
}

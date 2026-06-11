package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) Lookup(
	ctx context.Context,
	query ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	_ = ctx
	_ = query

	return nil, errProductRepositoryNotImplemented
}

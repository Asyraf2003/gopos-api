package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) List(
	ctx context.Context,
	query ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	_ = ctx
	_ = query

	return nil, errProductRepositoryNotImplemented
}

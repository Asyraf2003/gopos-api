package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) Append(
	ctx context.Context,
	version ports.ProductVersionRecord,
) error {
	_ = ctx
	_ = version

	return errProductRepositoryNotImplemented
}

func (r *ProductRepository) ListByProductID(
	ctx context.Context,
	productID string,
) ([]ports.ProductVersionRecord, error) {
	_ = ctx
	_ = productID

	return nil, errProductRepositoryNotImplemented
}

package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) CheckCreateDuplicate(
	ctx context.Context,
	candidate ports.ProductDuplicateCandidate,
) error {
	_ = ctx
	_ = candidate

	return errProductRepositoryNotImplemented
}

func (r *ProductRepository) CheckUpdateDuplicate(
	ctx context.Context,
	productID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	_ = ctx
	_ = productID
	_ = candidate

	return errProductRepositoryNotImplemented
}

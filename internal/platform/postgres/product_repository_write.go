package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
)

func (r *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	_ = ctx
	_ = product

	return errProductRepositoryNotImplemented
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	_ = ctx
	_ = product

	return errProductRepositoryNotImplemented
}

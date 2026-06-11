package postgres

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
)

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	_ = ctx
	_ = id

	return nil, errProductRepositoryNotImplemented
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	_ = ctx
	_ = id

	return nil, errProductRepositoryNotImplemented
}

package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
)

type softDeleteProductRepositoryDouble struct {
	found   *domain.Product
	updated *domain.Product
	err     error
	findErr error
}

func (d *softDeleteProductRepositoryDouble) Create(
	_ context.Context,
	_ *domain.Product,
) error {
	return nil
}

func (d *softDeleteProductRepositoryDouble) Update(
	_ context.Context,
	product *domain.Product,
) error {
	d.updated = product
	return d.err
}

func (d *softDeleteProductRepositoryDouble) FindByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	if d.findErr != nil {
		return nil, d.findErr
	}

	return d.found, nil
}

package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type fakeProductIDGenerator struct {
	id  string
	err error
}

func (f fakeProductIDGenerator) NewProductID() (string, error) {
	if f.err != nil {
		return "", f.err
	}

	return f.id, nil
}

type fakeProductRepository struct {
	created *domain.Product
	err     error
}

func (f *fakeProductRepository) Create(_ context.Context, product *domain.Product) error {
	f.created = product

	return f.err
}

func (f *fakeProductRepository) Update(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeProductRepository) FindByID(_ context.Context, _ string) (*domain.Product, error) {
	return nil, ports.ErrProductNotFound
}

type fakeProductDuplicateChecker struct {
	createCalled bool
	candidate    ports.ProductDuplicateCandidate
	err          error
}

func (f *fakeProductDuplicateChecker) CheckCreateDuplicate(
	_ context.Context,
	candidate ports.ProductDuplicateCandidate,
) error {
	f.createCalled = true
	f.candidate = candidate

	return f.err
}

func (f *fakeProductDuplicateChecker) CheckUpdateDuplicate(
	_ context.Context,
	_ string,
	_ ports.ProductDuplicateCandidate,
) error {
	return nil
}

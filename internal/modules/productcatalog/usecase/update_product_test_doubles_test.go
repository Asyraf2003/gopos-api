package usecase

import (
	"context"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"
)

type fakeUpdateProductRepository struct {
	product *domain.Product
}

func (f *fakeUpdateProductRepository) Create(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeUpdateProductRepository) Update(_ context.Context, _ *domain.Product) error {
	return nil
}

func (f *fakeUpdateProductRepository) FindByID(
	_ context.Context,
	_ string,
) (*domain.Product, error) {
	return f.product, nil
}

type fakeUpdateProductDuplicateChecker struct {
	updateCalled bool
	productID    string
	candidate    ports.ProductDuplicateCandidate
	err          error
}

func (f *fakeUpdateProductDuplicateChecker) CheckCreateDuplicate(
	_ context.Context,
	_ ports.ProductDuplicateCandidate,
) error {
	return nil
}

func (f *fakeUpdateProductDuplicateChecker) CheckUpdateDuplicate(
	_ context.Context,
	productID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	f.updateCalled = true
	f.productID = productID
	f.candidate = candidate

	return f.err
}

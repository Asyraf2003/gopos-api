package usecase

import (
	"context"

	"pos-go/internal/modules/servicecatalog/domain"
)

type fakeServiceCatalogRepository struct {
	items map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem
}

func newFakeServiceCatalogRepository() *fakeServiceCatalogRepository {
	return &fakeServiceCatalogRepository{
		items: make(map[domain.ServiceCatalogItemID]domain.ServiceCatalogItem),
	}
}

func (r *fakeServiceCatalogRepository) Create(
	_ context.Context,
	item domain.ServiceCatalogItem,
) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) Update(
	_ context.Context,
	item domain.ServiceCatalogItem,
) error {
	r.items[item.ID()] = item
	return nil
}

func (r *fakeServiceCatalogRepository) FindByID(
	_ context.Context,
	id domain.ServiceCatalogItemID,
) (domain.ServiceCatalogItem, bool, error) {
	item, found := r.items[id]
	return item, found, nil
}

func (r *fakeServiceCatalogRepository) FindByNormalizedName(
	_ context.Context,
	normalizedName domain.NormalizedName,
) (domain.ServiceCatalogItem, bool, error) {
	for _, item := range r.items {
		if item.NormalizedName() == normalizedName {
			return item, true, nil
		}
	}

	return domain.ServiceCatalogItem{}, false, nil
}

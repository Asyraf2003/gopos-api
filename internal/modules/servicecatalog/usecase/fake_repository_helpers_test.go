package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/servicecatalog/domain"
)

func seedServiceCatalogItem(
	t *testing.T,
	repo *fakeServiceCatalogRepository,
	id domain.ServiceCatalogItemID,
	name string,
	price domain.MoneyRupiah,
) domain.ServiceCatalogItem {
	t.Helper()

	item, err := domain.NewServiceCatalogItem(id, name, price, fixedNow())
	if err != nil {
		t.Fatalf("NewServiceCatalogItem() error = %v", err)
	}

	if err := repo.Create(context.Background(), item); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	return item
}

func fixedIDGenerator(id domain.ServiceCatalogItemID) ServiceCatalogItemIDGenerator {
	return func() (domain.ServiceCatalogItemID, error) {
		return id, nil
	}
}

func fixedClock() time.Time {
	return time.Date(2026, 6, 8, 10, 0, 0, 0, time.UTC)
}

func fixedNow() time.Time {
	return fixedClock()
}

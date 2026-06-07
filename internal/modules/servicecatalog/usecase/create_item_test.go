package usecase

import (
	"context"
	"errors"
	"testing"
)

func TestCreateServiceCatalogItemStoresItem(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	uc := NewCreateServiceCatalogItem(repo, fixedIDGenerator("svc_1"), fixedClock)

	got, err := uc.Execute(ctx, CreateServiceCatalogItemCommand{
		Name:               "  Potong Rambut  ",
		DefaultPriceRupiah: 10000,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if got.ID != "svc_1" {
		t.Fatalf("ID = %q, want %q", got.ID, "svc_1")
	}

	if !got.IsActive {
		t.Fatal("created item should be active by default")
	}

	stored, found, err := repo.FindByID(ctx, "svc_1")
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}

	if !found {
		t.Fatal("created item was not stored")
	}

	if stored.Name() != "Potong Rambut" {
		t.Fatalf("stored Name() = %q, want %q", stored.Name(), "Potong Rambut")
	}
}

func TestCreateServiceCatalogItemRejectsDuplicateNormalizedName(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_existing", "Potong Rambut", 10000)

	uc := NewCreateServiceCatalogItem(repo, fixedIDGenerator("svc_new"), fixedClock)

	_, err := uc.Execute(ctx, CreateServiceCatalogItemCommand{
		Name:               "potong   rambut",
		DefaultPriceRupiah: 12000,
	})
	if !errors.Is(err, ErrDuplicateServiceCatalogItemNormalizedName) {
		t.Fatalf("error = %v, want duplicate normalized name", err)
	}
}

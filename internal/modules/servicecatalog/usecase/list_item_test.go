package usecase

import (
	"context"
	"testing"

	"pos-go/internal/modules/servicecatalog/ports"
)

func TestListServiceCatalogItemsFiltersActiveInactiveAndAll(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Potong Rambut", 10000)
	seedServiceCatalogItem(t, repo, "svc_2", "Cuci Motor", 15000)

	if _, _, err := repo.SetActive(ctx, "svc_2", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewListServiceCatalogItems(repo)

	active, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{
		Status: ports.ListStatusActive,
	})
	if err != nil {
		t.Fatalf("active Execute() error = %v", err)
	}

	if len(active) != 1 || active[0].ID != "svc_1" {
		t.Fatalf("active result = %+v, want only svc_1", active)
	}

	inactive, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{
		Status: ports.ListStatusInactive,
	})
	if err != nil {
		t.Fatalf("inactive Execute() error = %v", err)
	}

	if len(inactive) != 1 || inactive[0].ID != "svc_2" {
		t.Fatalf("inactive result = %+v, want only svc_2", inactive)
	}

	all, err := uc.Execute(ctx, ListServiceCatalogItemsCommand{
		Status: ports.ListStatusAll,
	})
	if err != nil {
		t.Fatalf("all Execute() error = %v", err)
	}

	if len(all) != 2 {
		t.Fatalf("all result count = %d, want %d", len(all), 2)
	}
}

package usecase

import (
	"context"
	"errors"
	"testing"
)

func TestLookupServiceCatalogItemsExcludesInactiveByDefault(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	seedServiceCatalogItem(t, repo, "svc_1", "Cuci Motor", 15000)
	seedServiceCatalogItem(t, repo, "svc_2", "Cuci Mobil", 25000)

	if _, _, err := repo.SetActive(ctx, "svc_2", false); err != nil {
		t.Fatalf("SetActive() error = %v", err)
	}

	uc := NewLookupServiceCatalogItems(repo)

	got, err := uc.Execute(ctx, LookupServiceCatalogItemsCommand{
		Query: "cuci",
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if len(got) != 1 || got[0].ID != "svc_1" {
		t.Fatalf("lookup result = %+v, want only active svc_1", got)
	}
}

func TestLookupServiceCatalogItemsEnforcesMaxLimit(t *testing.T) {
	ctx := context.Background()
	repo := newFakeServiceCatalogRepository()
	uc := NewLookupServiceCatalogItems(repo)

	_, err := uc.Execute(ctx, LookupServiceCatalogItemsCommand{
		Limit: 51,
	})
	if !errors.Is(err, ErrInvalidLookupLimit) {
		t.Fatalf("error = %v, want invalid lookup limit", err)
	}
}

// Copyright (C) 2026 Asyraf Mubarak
//
// This file is part of gopos-api.
//
// gopos-api is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, version 3 only.
//
// gopos-api is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with gopos-api. If not, see <https://www.gnu.org/licenses/>.

//go:build integration

package postgres

import (
	"context"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestProductRepository_LookupExcludesDeletedByDefault(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	active := newProductCatalogTestProduct(t, "Kampas Rem Lookup Active")
	deleted := newProductCatalogTestProduct(t, "Kampas Rem Lookup Deleted")
	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}
	if err := repo.Create(txCtx, deleted); err != nil {
		t.Fatalf("Create() deleted error = %v", err)
	}
	softDeleteProductForListTest(t, repo, txCtx, deleted)

	items, err := repo.Lookup(txCtx, ports.ProductLookupQuery{
		Query: "kampas rem lookup",
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("Lookup() len = %d, want 1", len(items))
	}
	if items[0].ID != active.ID() {
		t.Fatalf("Lookup() ID = %q, want %q", items[0].ID, active.ID())
	}
	if items[0].Status != "active" {
		t.Fatalf("Lookup() status = %q, want active", items[0].Status)
	}
}

func TestProductRepository_LookupIncludesDeletedWhenRequested(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	active := newProductCatalogTestProduct(t, "Filter Oli Lookup Active")
	deleted := newProductCatalogTestProduct(t, "Filter Oli Lookup Deleted")
	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}
	if err := repo.Create(txCtx, deleted); err != nil {
		t.Fatalf("Create() deleted error = %v", err)
	}
	softDeleteProductForListTest(t, repo, txCtx, deleted)

	items, err := repo.Lookup(txCtx, ports.ProductLookupQuery{
		Query:          "filter oli lookup",
		Limit:          10,
		IncludeDeleted: true,
	})
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("Lookup() len = %d, want 2", len(items))
	}
}

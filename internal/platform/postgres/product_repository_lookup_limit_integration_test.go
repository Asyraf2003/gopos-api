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

func TestProductRepository_LookupRespectsLimit(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	products := []string{
		"Aki Lookup Limit",
		"Ban Lookup Limit",
		"Busi Lookup Limit",
	}
	for _, name := range products {
		product := newProductCatalogTestProduct(t, name)
		if err := repo.Create(txCtx, product); err != nil {
			t.Fatalf("Create() error = %v", err)
		}
	}

	items, err := repo.Lookup(txCtx, ports.ProductLookupQuery{
		Query: "lookup limit",
		Limit: 2,
	})
	if err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("Lookup() len = %d, want 2", len(items))
	}
	if items[0].Name != "Aki Lookup Limit" {
		t.Fatalf("Lookup() first name = %q, want Aki Lookup Limit", items[0].Name)
	}
	if items[1].Name != "Ban Lookup Limit" {
		t.Fatalf("Lookup() second name = %q, want Ban Lookup Limit", items[1].Name)
	}
}

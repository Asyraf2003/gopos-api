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
	"errors"
	"testing"

	"pos-go/internal/modules/productcatalog/ports"
)

func TestProductRepository_CreateAndFindByID(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	product := newProductCatalogTestProduct(t, "Kampas Rem")

	if err := repo.Create(txCtx, product); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	got, err := repo.FindByID(txCtx, product.ID())
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}

	if got.ID() != product.ID() {
		t.Fatalf("ID() = %q, want %q", got.ID(), product.ID())
	}
	if got.Code() == nil || product.Code() == nil || *got.Code() != *product.Code() {
		t.Fatalf("Code() = %v, want %v", got.Code(), product.Code())
	}
	if got.Name() != product.Name() {
		t.Fatalf("Name() = %q, want %q", got.Name(), product.Name())
	}
	if got.NormalizedName() != product.NormalizedName() {
		t.Fatalf("NormalizedName() = %q, want %q", got.NormalizedName(), product.NormalizedName())
	}
	if got.Brand() != product.Brand() {
		t.Fatalf("Brand() = %q, want %q", got.Brand(), product.Brand())
	}
	if got.NormalizedBrand() != product.NormalizedBrand() {
		t.Fatalf("NormalizedBrand() = %q, want %q", got.NormalizedBrand(), product.NormalizedBrand())
	}
	if got.SalePriceRupiah() != product.SalePriceRupiah() {
		t.Fatalf("SalePriceRupiah() = %d, want %d", got.SalePriceRupiah(), product.SalePriceRupiah())
	}
}

func TestProductRepository_FindByIDMissing(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	_, err := repo.FindByID(txCtx, "missing-product")
	if !errors.Is(err, ports.ErrProductNotFound) {
		t.Fatalf("FindByID() error = %v, want %v", err, ports.ErrProductNotFound)
	}
}

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

	"pos-go/internal/modules/productcatalog/domain"
)

func TestProductRepository_Update(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	product := newProductCatalogTestProduct(t, "Kampas Rem")
	if err := repo.Create(txCtx, product); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	updated, err := domain.NewProduct(domain.ProductInput{
		ID:              product.ID(),
		Code:            "SKU-UPDATED",
		Name:            "Kampas Rem Racing",
		Brand:           "Yamaha",
		Size:            domain.IntPtr(120),
		SalePriceRupiah: 65000,
	})
	if err != nil {
		t.Fatalf("NewProduct() updated error = %v", err)
	}

	if err := repo.Update(txCtx, updated); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := repo.FindByID(txCtx, product.ID())
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}

	if got.ID() != updated.ID() {
		t.Fatalf("ID() = %q, want %q", got.ID(), updated.ID())
	}
	if got.Code() == nil || *got.Code() != "SKU-UPDATED" {
		t.Fatalf("Code() = %v, want SKU-UPDATED", got.Code())
	}
	if got.Name() != "Kampas Rem Racing" {
		t.Fatalf("Name() = %q, want Kampas Rem Racing", got.Name())
	}
	if got.NormalizedName() != "kampas rem racing" {
		t.Fatalf("NormalizedName() = %q, want kampas rem racing", got.NormalizedName())
	}
	if got.Brand() != "Yamaha" {
		t.Fatalf("Brand() = %q, want Yamaha", got.Brand())
	}
	if got.NormalizedBrand() != "yamaha" {
		t.Fatalf("NormalizedBrand() = %q, want yamaha", got.NormalizedBrand())
	}
	if got.Size() == nil || *got.Size() != 120 {
		t.Fatalf("Size() = %v, want 120", got.Size())
	}
	if got.SalePriceRupiah() != 65000 {
		t.Fatalf("SalePriceRupiah() = %d, want 65000", got.SalePriceRupiah())
	}
}

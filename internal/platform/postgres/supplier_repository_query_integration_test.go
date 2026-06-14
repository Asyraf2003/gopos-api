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

	"pos-go/internal/modules/supplier/domain"
)

func TestSupplierRepository_FindByIDMissing(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)

	_, found, err := repo.FindByID(txCtx, domain.SupplierID("missing-supplier"))
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if found {
		t.Fatal("FindByID() found = true, want false")
	}
}

func TestSupplierRepository_FindByNormalizedName(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Sentosa Parts")

	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	got, found, err := repo.FindByNormalizedName(txCtx, domain.NormalizeName("sentosa parts"))
	if err != nil {
		t.Fatalf("FindByNormalizedName() error = %v", err)
	}
	if !found {
		t.Fatal("FindByNormalizedName() found = false, want true")
	}
	assertSupplierFields(t, got, supplier)
}

func TestSupplierRepository_FindActiveByNormalizedName(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	inactive := newInactiveSupplierRepositoryTestSupplier(t, "Terang Abadi")
	active := newSupplierRepositoryTestSupplier(t, " terang  abadi ")

	if err := repo.Create(txCtx, inactive); err != nil {
		t.Fatalf("Create() inactive error = %v", err)
	}
	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active error = %v", err)
	}

	got, found, err := repo.FindActiveByNormalizedName(txCtx, active.NormalizedName())
	if err != nil {
		t.Fatalf("FindActiveByNormalizedName() error = %v", err)
	}
	if !found {
		t.Fatal("FindActiveByNormalizedName() found = false, want true")
	}
	if got.ID() != active.ID() {
		t.Fatalf("ID() = %q, want active %q", got.ID(), active.ID())
	}
}

func TestSupplierRepository_FindActiveByNormalizedNameIgnoresInactive(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	inactive := newInactiveSupplierRepositoryTestSupplier(t, "Cahaya Motor")

	if err := repo.Create(txCtx, inactive); err != nil {
		t.Fatalf("Create() inactive error = %v", err)
	}

	_, found, err := repo.FindActiveByNormalizedName(txCtx, inactive.NormalizedName())
	if err != nil {
		t.Fatalf("FindActiveByNormalizedName() error = %v", err)
	}
	if found {
		t.Fatal("FindActiveByNormalizedName() found = true, want false")
	}
}

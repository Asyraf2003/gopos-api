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

	"github.com/jackc/pgx/v5/pgconn"
)

func TestSupplierRepository_CreateStoresFields(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	supplier := newSupplierRepositoryTestSupplier(t, "Bengkel Jaya")

	if err := repo.Create(txCtx, supplier); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	got, found, err := repo.FindByID(txCtx, supplier.ID())
	if err != nil {
		t.Fatalf("FindByID() error = %v", err)
	}
	if !found {
		t.Fatal("FindByID() found = false, want true")
	}
	assertSupplierFields(t, got, supplier)

	stored := mustReadSupplierNormalizedName(t, txCtx, repo, supplier.ID())
	if stored != string(supplier.NormalizedName()) {
		t.Fatalf("stored normalized name = %q, want %q", stored, supplier.NormalizedName())
	}
}

func TestSupplierRepository_CreateRejectsDuplicateActiveNormalizedName(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)

	if err := repo.Create(txCtx, newSupplierRepositoryTestSupplier(t, "Mitra Parts")); err != nil {
		t.Fatalf("Create() existing error = %v", err)
	}

	err := repo.Create(txCtx, newSupplierRepositoryTestSupplier(t, " mitra   parts "))
	if !isSupplierActiveNameUniqueViolation(err) {
		t.Fatalf("Create() error = %v, want active name unique violation", err)
	}
}

func TestSupplierRepository_CreateAllowsInactiveNameReuse(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)

	inactive := newInactiveSupplierRepositoryTestSupplier(t, "Sinar Motor")
	if err := repo.Create(txCtx, inactive); err != nil {
		t.Fatalf("Create() inactive error = %v", err)
	}

	active := newSupplierRepositoryTestSupplier(t, " sinar   motor ")
	if err := repo.Create(txCtx, active); err != nil {
		t.Fatalf("Create() active reused name error = %v", err)
	}
}

func isSupplierActiveNameUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) &&
		pgErr.Code == "23505" &&
		pgErr.ConstraintName == "suppliers_active_name_normalized_unique"
}

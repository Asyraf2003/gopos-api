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
	"time"

	"pos-go/internal/modules/supplier/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var supplierRepositoryTestTime = time.Date(2026, 6, 14, 8, 30, 0, 123456000, time.UTC)

func mustOpenSupplierRepoTx(t *testing.T, ctx context.Context) (*pgxpool.Pool, context.Context) {
	t.Helper()

	pool := mustOpenIntegrationPool(t, ctx)
	tx := mustBeginIntegrationTx(t, ctx, pool)
	t.Cleanup(func() {
		_ = tx.Rollback(ctx)
		pool.Close()
	})

	return pool, contextWithTx(ctx, tx)
}

func newSupplierRepositoryTestSupplier(t *testing.T, name string) domain.Supplier {
	t.Helper()

	supplier, err := domain.NewSupplier(
		domain.SupplierID(uuid.NewString()),
		name,
		domain.SupplierContact{
			Phone:   "08123456789",
			Email:   "supplier@example.com",
			Address: "Jalan Supplier 1",
			Notes:   "primary supplier",
		},
		supplierRepositoryTestTime,
	)
	if err != nil {
		t.Fatalf("NewSupplier() error = %v", err)
	}

	return supplier
}

func newInactiveSupplierRepositoryTestSupplier(t *testing.T, name string) domain.Supplier {
	t.Helper()

	supplier := newSupplierRepositoryTestSupplier(t, name)
	supplier.Deactivate(supplierRepositoryTestTime.Add(time.Minute))
	return supplier
}

func mustReadSupplierNormalizedName(
	t *testing.T,
	ctx context.Context,
	repo *SupplierRepository,
	id domain.SupplierID,
) string {
	t.Helper()

	var normalizedName string
	err := repo.queryRow(ctx, `
		SELECT name_normalized
		FROM suppliers
		WHERE id = $1
	`, string(id)).Scan(&normalizedName)
	if err != nil {
		t.Fatalf("read supplier normalized name error = %v", err)
	}

	return normalizedName
}

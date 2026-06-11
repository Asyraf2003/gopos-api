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

func TestProductDuplicateChecker_CheckCreateDuplicateRejectsActiveCode(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	existing := newDuplicateCheckerProduct(t, "SKU-DUP-CODE-A", "Kampas Rem Dup Code")
	if err := repo.Create(txCtx, existing); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	candidate := newDuplicateCheckerCandidate(t, "SKU-DUP-CODE-A", "Busi Dup Code")
	err := repo.CheckCreateDuplicate(txCtx, candidate)
	if !errors.Is(err, ports.ErrDuplicateProductCode) {
		t.Fatalf("CheckCreateDuplicate() error = %v, want duplicate code", err)
	}
}

func TestProductDuplicateChecker_CheckCreateDuplicateRejectsIdentityWithoutDistinctCodes(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	existing := newDuplicateCheckerProduct(t, "", "Kampas Rem Dup Identity")
	if err := repo.Create(txCtx, existing); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	candidate := newDuplicateCheckerCandidate(t, "", "Kampas Rem Dup Identity")
	err := repo.CheckCreateDuplicate(txCtx, candidate)
	if !errors.Is(err, ports.ErrDuplicateProductIdentity) {
		t.Fatalf("CheckCreateDuplicate() error = %v, want duplicate identity", err)
	}
}

func TestProductDuplicateChecker_CheckCreateDuplicateAllowsDistinctCodedIdentity(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	existing := newDuplicateCheckerProduct(t, "SKU-DUP-IDENTITY-A", "Filter Oli Dup Allowed")
	if err := repo.Create(txCtx, existing); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	candidate := newDuplicateCheckerCandidate(t, "SKU-DUP-IDENTITY-B", "Filter Oli Dup Allowed")
	if err := repo.CheckCreateDuplicate(txCtx, candidate); err != nil {
		t.Fatalf("CheckCreateDuplicate() error = %v", err)
	}
}

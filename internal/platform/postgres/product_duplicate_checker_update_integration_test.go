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

func TestProductDuplicateChecker_CheckUpdateDuplicateIgnoresSameProduct(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	product := newDuplicateCheckerProduct(t, "SKU-DUP-SELF", "Aki Dup Self")
	if err := repo.Create(txCtx, product); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	candidate := duplicateCandidateFromProduct(product)
	if err := repo.CheckUpdateDuplicate(txCtx, product.ID(), candidate); err != nil {
		t.Fatalf("CheckUpdateDuplicate() error = %v", err)
	}
}

func TestProductDuplicateChecker_CheckUpdateDuplicateRejectsOtherProductCode(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	existing := newDuplicateCheckerProduct(t, "SKU-DUP-OTHER", "Aki Dup Other")
	updated := newDuplicateCheckerProduct(t, "SKU-DUP-UPDATING", "Ban Dup Other")
	if err := repo.Create(txCtx, existing); err != nil {
		t.Fatalf("Create() existing error = %v", err)
	}
	if err := repo.Create(txCtx, updated); err != nil {
		t.Fatalf("Create() updated error = %v", err)
	}

	candidate := newDuplicateCheckerCandidate(t, "SKU-DUP-OTHER", "Ban Dup Other")
	err := repo.CheckUpdateDuplicate(txCtx, updated.ID(), candidate)
	if !errors.Is(err, ports.ErrDuplicateProductCode) {
		t.Fatalf("CheckUpdateDuplicate() error = %v, want duplicate code", err)
	}
}

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
)

func TestProductDuplicateChecker_CheckCreateDuplicateIgnoresDeleted(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenProductCatalogRepoTx(t, ctx)
	repo := NewProductRepository(pool)

	deleted := newDuplicateCheckerProduct(t, "SKU-DUP-DELETED", "Busi Dup Deleted")
	if err := repo.Create(txCtx, deleted); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	softDeleteProductForListTest(t, repo, txCtx, deleted)

	candidate := newDuplicateCheckerCandidate(t, "SKU-DUP-DELETED", "Busi Dup Deleted")
	if err := repo.CheckCreateDuplicate(txCtx, candidate); err != nil {
		t.Fatalf("CheckCreateDuplicate() error = %v", err)
	}
}

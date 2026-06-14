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
	"fmt"
	"testing"

	"pos-go/internal/modules/supplier/ports"
)

func TestSupplierRepository_ListAndLookupDirectCallHardening(t *testing.T) {
	ctx := context.Background()
	pool, txCtx := mustOpenSupplierRepoTx(t, ctx)
	repo := NewSupplierRepository(pool)
	for _, row := range []struct {
		name   string
		active bool
	}{
		{"Alpha   Tools", true},
		{"Beta Tools", false},
		{"Gamma Tools", true},
		{"Omega Supplies", true},
	} {
		mustCreateSupplierQueryRow(t, txCtx, repo, row.name, row.active)
	}
	for i := range 55 {
		mustCreateSupplierQueryRow(t, txCtx, repo, fmt.Sprintf("ZZ Cap %02d", i), true)
	}

	list := func(filter ports.ListSuppliersFilter, names ...string) {
		rows, err := repo.List(txCtx, filter)
		assertSupplierNames(t, rows, err, names...)
	}
	lookup := func(filter ports.LookupSuppliersFilter, names ...string) {
		rows, err := repo.Lookup(txCtx, filter)
		assertSupplierNames(t, rows, err, names...)
	}
	listLen := func(filter ports.ListSuppliersFilter, want int) {
		rows, err := repo.List(txCtx, filter)
		if err != nil || len(rows) != want {
			t.Fatalf("List() len = %d err %v, want %d nil", len(rows), err, want)
		}
	}
	lookupLen := func(filter ports.LookupSuppliersFilter, want int) {
		rows, err := repo.Lookup(txCtx, filter)
		if err != nil || len(rows) != want {
			t.Fatalf("Lookup() len = %d err %v, want %d nil", len(rows), err, want)
		}
	}

	list(ports.ListSuppliersFilter{Page: 0, PerPage: 0}, "Alpha Tools", "Gamma Tools", "Omega Supplies", "ZZ Cap 00", "ZZ Cap 01", "ZZ Cap 02", "ZZ Cap 03", "ZZ Cap 04", "ZZ Cap 05", "ZZ Cap 06")
	list(ports.ListSuppliersFilter{Status: ports.ListStatusFilter("weird"), Page: -9, PerPage: 3}, "Alpha Tools", "Gamma Tools", "Omega Supplies")
	list(ports.ListSuppliersFilter{Query: "   ", Status: ports.ListStatusAll, Page: 1, PerPage: 4}, "Alpha Tools", "Beta Tools", "Gamma Tools", "Omega Supplies")
	list(ports.ListSuppliersFilter{Query: "ALPHA TOOLS", Status: ports.ListStatusAll, Page: 1, PerPage: 10}, "Alpha Tools")
	list(ports.ListSuppliersFilter{Query: "alpha    tools", Status: ports.ListStatusAll, Page: 1, PerPage: 10}, "Alpha Tools")
	listLen(ports.ListSuppliersFilter{Status: ports.ListStatusAll, Page: 1, PerPage: 99}, 50)

	lookup(ports.LookupSuppliersFilter{Limit: 0, ActiveOnly: true}, "Alpha Tools", "Gamma Tools", "Omega Supplies", "ZZ Cap 00", "ZZ Cap 01", "ZZ Cap 02", "ZZ Cap 03", "ZZ Cap 04", "ZZ Cap 05", "ZZ Cap 06", "ZZ Cap 07", "ZZ Cap 08", "ZZ Cap 09", "ZZ Cap 10", "ZZ Cap 11", "ZZ Cap 12", "ZZ Cap 13", "ZZ Cap 14", "ZZ Cap 15", "ZZ Cap 16")
	lookup(ports.LookupSuppliersFilter{Query: "   ", Limit: 4}, "Alpha Tools", "Beta Tools", "Gamma Tools", "Omega Supplies")
	lookup(ports.LookupSuppliersFilter{Query: "BETA TOOLS", Limit: 10}, "Beta Tools")
	lookup(ports.LookupSuppliersFilter{Query: "alpha    tools", Limit: 10}, "Alpha Tools")
	lookupLen(ports.LookupSuppliersFilter{Limit: 99, ActiveOnly: true}, 50)
}

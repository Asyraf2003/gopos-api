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
// along with gopos-api. If not, see https://www.gnu.org/licenses/.

package usecase

import (
	"context"
	"testing"
	"time"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

type fakeSupplierRepository struct {
	byID map[domain.SupplierID]domain.Supplier

	createCalls  int
	updateCalls  int
	listFilter   ports.ListSuppliersFilter
	lookupFilter ports.LookupSuppliersFilter
}

func newFakeSupplierRepository() *fakeSupplierRepository {
	return &fakeSupplierRepository{byID: map[domain.SupplierID]domain.Supplier{}}
}

func (r *fakeSupplierRepository) Create(_ context.Context, supplier domain.Supplier) error {
	r.createCalls++
	r.byID[supplier.ID()] = supplier
	return nil
}

func (r *fakeSupplierRepository) Update(_ context.Context, supplier domain.Supplier) error {
	r.updateCalls++
	r.byID[supplier.ID()] = supplier
	return nil
}

func (r *fakeSupplierRepository) FindByID(_ context.Context, id domain.SupplierID) (domain.Supplier, bool, error) {
	supplier, found := r.byID[id]
	return supplier, found, nil
}

func (r *fakeSupplierRepository) FindByNormalizedName(
	_ context.Context,
	normalizedName domain.NormalizedName,
) (domain.Supplier, bool, error) {
	for _, supplier := range r.byID {
		if supplier.NormalizedName() == normalizedName {
			return supplier, true, nil
		}
	}

	return domain.Supplier{}, false, nil
}

func (r *fakeSupplierRepository) List(
	_ context.Context,
	filter ports.ListSuppliersFilter,
) ([]domain.Supplier, error) {
	r.listFilter = filter
	results := make([]domain.Supplier, 0, len(r.byID))
	for _, supplier := range r.byID {
		results = append(results, supplier)
	}

	return results, nil
}

func (r *fakeSupplierRepository) Lookup(
	_ context.Context,
	filter ports.LookupSuppliersFilter,
) ([]domain.Supplier, error) {
	r.lookupFilter = filter
	results := make([]domain.Supplier, 0, len(r.byID))
	for _, supplier := range r.byID {
		results = append(results, supplier)
	}

	return results, nil
}

func (r *fakeSupplierRepository) SetActive(
	_ context.Context,
	id domain.SupplierID,
	active bool,
) (domain.Supplier, bool, error) {
	supplier, found := r.byID[id]
	if !found {
		return domain.Supplier{}, false, nil
	}

	if active {
		supplier.Activate(time.Date(2026, 6, 14, 12, 0, 0, 0, time.UTC))
	} else {
		supplier.Deactivate(time.Date(2026, 6, 14, 12, 0, 0, 0, time.UTC))
	}
	r.byID[id] = supplier

	return supplier, true, nil
}

func mustSupplier(t *testing.T, id string, name string, active bool) domain.Supplier {
	t.Helper()

	supplier, err := domain.NewSupplier(
		domain.SupplierID(id),
		name,
		domain.SupplierContact{Phone: "0812", Email: "owner@example.test"},
		time.Date(2026, 6, 14, 10, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatalf("NewSupplier() error = %v", err)
	}

	if !active {
		supplier.Deactivate(time.Date(2026, 6, 14, 11, 0, 0, 0, time.UTC))
	}

	return supplier
}

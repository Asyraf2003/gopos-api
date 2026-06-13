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

package usecase

import (
	"context"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"
)

type fakeSupplierRepository struct {
	byID         map[domain.SupplierID]domain.Supplier
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

func (r *fakeSupplierRepository) FindActiveByNormalizedName(
	_ context.Context,
	normalizedName domain.NormalizedName,
) (domain.Supplier, bool, error) {
	for _, supplier := range r.byID {
		if supplier.NormalizedName() == normalizedName && supplier.IsActive() {
			return supplier, true, nil
		}
	}

	return domain.Supplier{}, false, nil
}

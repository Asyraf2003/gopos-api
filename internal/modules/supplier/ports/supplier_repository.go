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

package ports

import (
	"context"

	"pos-go/internal/modules/supplier/domain"
)

type ListStatusFilter string

const (
	ListStatusActive   ListStatusFilter = "active"
	ListStatusInactive ListStatusFilter = "inactive"
	ListStatusAll      ListStatusFilter = "all"
)

type ListSuppliersFilter struct {
	Query   string
	Status  ListStatusFilter
	Page    int
	PerPage int
}

type LookupSuppliersFilter struct {
	Query      string
	Limit      int
	ActiveOnly bool
}

type SupplierRepository interface {
	Create(ctx context.Context, supplier domain.Supplier) error
	Update(ctx context.Context, supplier domain.Supplier) error
	FindByID(ctx context.Context, id domain.SupplierID) (domain.Supplier, bool, error)
	FindByNormalizedName(ctx context.Context, normalizedName domain.NormalizedName) (domain.Supplier, bool, error)
	FindActiveByNormalizedName(ctx context.Context, normalizedName domain.NormalizedName) (domain.Supplier, bool, error)
	List(ctx context.Context, filter ListSuppliersFilter) ([]domain.Supplier, error)
	Lookup(ctx context.Context, filter LookupSuppliersFilter) ([]domain.Supplier, error)
	SetActive(ctx context.Context, id domain.SupplierID, active bool) (domain.Supplier, bool, error)
}

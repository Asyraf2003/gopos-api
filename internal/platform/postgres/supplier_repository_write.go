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

package postgres

import (
	"context"

	"pos-go/internal/modules/supplier/domain"
)

func (r *SupplierRepository) Create(ctx context.Context, supplier domain.Supplier) error {
	sql := `
		INSERT INTO suppliers (
			id,
			name,
			name_normalized,
			phone,
			email,
			address,
			notes,
			is_active,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.exec(ctx, sql, supplierArgs(supplier)...)
	return err
}

func (r *SupplierRepository) Update(ctx context.Context, supplier domain.Supplier) error {
	return errSupplierRepositoryNotImplemented
}

func (r *SupplierRepository) SetActive(
	ctx context.Context,
	id domain.SupplierID,
	active bool,
) (domain.Supplier, bool, error) {
	return domain.Supplier{}, false, errSupplierRepositoryNotImplemented
}

func supplierArgs(supplier domain.Supplier) []any {
	return []any{
		string(supplier.ID()),
		supplier.Name(),
		string(supplier.NormalizedName()),
		supplierNullableText(supplier.Phone()),
		supplierNullableText(supplier.Email()),
		supplierNullableText(supplier.Address()),
		supplierNullableText(supplier.Notes()),
		supplier.IsActive(),
		supplier.CreatedAt(),
		supplier.UpdatedAt(),
	}
}

func supplierNullableText(value string) any {
	if value == "" {
		return nil
	}

	return value
}

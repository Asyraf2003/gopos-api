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

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) checkProductDuplicateCode(
	ctx context.Context,
	excludeProductID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	if candidate.Code == nil {
		return nil
	}

	exists, err := r.productDuplicateCodeExists(ctx, excludeProductID, *candidate.Code)
	if err != nil {
		return err
	}
	if exists {
		return ports.ErrDuplicateProductCode
	}

	return nil
}

func (r *ProductRepository) productDuplicateCodeExists(
	ctx context.Context,
	excludeProductID string,
	code string,
) (bool, error) {
	var exists bool
	err := r.queryRow(ctx,
		productDuplicateCodeSQL(excludeProductID),
		productDuplicateCodeArgs(excludeProductID, code)...,
	).Scan(&exists)

	return exists, err
}

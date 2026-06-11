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
	"database/sql"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) checkProductDuplicateIdentity(
	ctx context.Context,
	excludeProductID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	rows, err := r.query(ctx,
		productDuplicateIdentitySQL(excludeProductID),
		productDuplicateIdentityArgs(excludeProductID, candidate)...,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var existingCode sql.NullString
		if err := rows.Scan(&existingCode); err != nil {
			return err
		}
		if !productDuplicateIdentityAllowed(candidate.Code, existingCode) {
			return ports.ErrDuplicateProductIdentity
		}
	}

	return rows.Err()
}

func productDuplicateIdentityAllowed(candidateCode *string, existingCode sql.NullString) bool {
	if candidateCode == nil || !existingCode.Valid {
		return false
	}

	return existingCode.String != *candidateCode
}

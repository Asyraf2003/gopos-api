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
	"database/sql"

	"pos-go/internal/modules/productcatalog/ports"
)

func productDuplicateCodeSQL(excludeProductID string) string {
	if excludeProductID == "" {
		return `
			SELECT EXISTS (
				SELECT 1 FROM products
				WHERE deleted_at IS NULL AND kode_barang = $1
			)
		`
	}

	return `
		SELECT EXISTS (
			SELECT 1 FROM products
			WHERE deleted_at IS NULL AND kode_barang = $1 AND id <> $2
		)
	`
}

func productDuplicateCodeArgs(excludeProductID string, code string) []any {
	if excludeProductID == "" {
		return []any{code}
	}

	return []any{code, excludeProductID}
}

func productDuplicateIdentitySQL(excludeProductID string) string {
	if excludeProductID == "" {
		return productDuplicateIdentityBaseSQL("")
	}

	return productDuplicateIdentityBaseSQL("AND id <> $4")
}

func productDuplicateIdentityBaseSQL(excludeClause string) string {
	return `
		SELECT kode_barang
		FROM products
		WHERE deleted_at IS NULL
			AND nama_barang_normalized = $1
			AND merek_normalized = $2
			AND ukuran IS NOT DISTINCT FROM $3
			` + excludeClause + `
	`
}

func productDuplicateIdentityArgs(
	excludeProductID string,
	candidate ports.ProductDuplicateCandidate,
) []any {
	args := []any{
		candidate.NormalizedName,
		candidate.NormalizedBrand,
		nullableDuplicateSize(candidate.Size),
	}
	if excludeProductID != "" {
		args = append(args, excludeProductID)
	}

	return args
}

func nullableDuplicateSize(value *int) sql.NullInt64 {
	if value == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: int64(*value), Valid: true}
}

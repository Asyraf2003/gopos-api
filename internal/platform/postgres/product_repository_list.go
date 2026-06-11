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
	"fmt"
	"strings"

	"pos-go/internal/modules/productcatalog/ports"
)

const (
	productListDefaultPage    = 1
	productListDefaultPerPage = 20
	productListMaxPerPage     = 100
)

func (r *ProductRepository) List(
	ctx context.Context,
	query ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	where, args := productListWhere(query)

	perPage := query.PerPage
	if perPage <= 0 {
		perPage = productListDefaultPerPage
	}
	if perPage > productListMaxPerPage {
		perPage = productListMaxPerPage
	}

	page := query.Page
	if page <= 0 {
		page = productListDefaultPage
	}

	args = append(args, perPage)
	limitPlaceholder := fmt.Sprintf("$%d", len(args))

	args = append(args, (page-1)*perPage)
	offsetPlaceholder := fmt.Sprintf("$%d", len(args))

	querySQL := `
		SELECT
			id,
			kode_barang,
			nama_barang,
			merek,
			ukuran,
			harga_jual,
			CASE
				WHEN deleted_at IS NULL THEN 'active'
				ELSE 'deleted'
			END AS status
		FROM products
	`
	if len(where) > 0 {
		querySQL += "\n\t\tWHERE " + strings.Join(where, "\n\t\t\tAND ")
	}
	querySQL += fmt.Sprintf(`
		ORDER BY
			nama_barang_normalized ASC,
			merek_normalized ASC,
			ukuran ASC NULLS LAST,
			id ASC
		LIMIT %s
		OFFSET %s
	`, limitPlaceholder, offsetPlaceholder)

	rows, err := r.query(ctx, querySQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []ports.ProductListItem{}
	for rows.Next() {
		item, err := scanProductListItem(rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

func productListWhere(query ports.ProductListQuery) ([]string, []any) {
	where := []string{}
	args := []any{}

	switch strings.ToLower(strings.TrimSpace(query.Status)) {
	case "deleted":
		where = append(where, "deleted_at IS NOT NULL")
	case "all":
	default:
		where = append(where, "deleted_at IS NULL")
	}

	search := strings.TrimSpace(query.Search)
	if search != "" {
		args = append(args, "%"+strings.ToLower(search)+"%")
		placeholder := fmt.Sprintf("$%d", len(args))
		where = append(where, fmt.Sprintf(
			"(nama_barang_normalized LIKE %[1]s OR merek_normalized LIKE %[1]s OR kode_barang ILIKE %[1]s)",
			placeholder,
		))
	}

	return where, args
}

func scanProductListItem(row productScanner) (ports.ProductListItem, error) {
	var item ports.ProductListItem
	var code sql.NullString
	var size sql.NullInt64

	err := row.Scan(
		&item.ID,
		&code,
		&item.Name,
		&item.Brand,
		&size,
		&item.SalePriceRupiah,
		&item.Status,
	)
	if err != nil {
		return ports.ProductListItem{}, err
	}

	item.Code = nullableStringPtr(code)
	item.Size = nullableIntPtr(size)

	return item, nil
}

func nullableStringPtr(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func nullableIntPtr(value sql.NullInt64) *int {
	if !value.Valid {
		return nil
	}

	intValue := int(value.Int64)
	return &intValue
}

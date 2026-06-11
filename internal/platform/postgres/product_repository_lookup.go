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

	"github.com/jackc/pgx/v5"
)

const (
	productLookupDefaultLimit = 10
	productLookupMaxLimit     = 50
)

func (r *ProductRepository) Lookup(
	ctx context.Context,
	query ports.ProductLookupQuery,
) ([]ports.ProductLookupItem, error) {
	where, args := productLookupWhere(query)
	limit := productLookupLimit(query)

	args = append(args, limit)
	limitPlaceholder := fmt.Sprintf("$%d", len(args))

	querySQL := productLookupSelectSQL()
	if len(where) > 0 {
		querySQL += "\n\t\tWHERE " + strings.Join(where, "\n\t\t\tAND ")
	}
	querySQL += fmt.Sprintf(productLookupOrderLimitSQL(), limitPlaceholder)

	rows, err := r.query(ctx, querySQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanProductLookupRows(rows)
}

func productLookupLimit(query ports.ProductLookupQuery) int {
	limit := query.Limit
	if limit <= 0 {
		limit = productLookupDefaultLimit
	}
	if limit > productLookupMaxLimit {
		limit = productLookupMaxLimit
	}

	return limit
}

func productLookupWhere(query ports.ProductLookupQuery) ([]string, []any) {
	where := []string{}
	args := []any{}

	if !query.IncludeDeleted {
		where = append(where, "deleted_at IS NULL")
	}

	search := strings.TrimSpace(query.Query)
	if search == "" {
		return where, args
	}

	args = append(args, "%"+strings.ToLower(search)+"%")
	placeholder := fmt.Sprintf("$%d", len(args))
	where = append(where, fmt.Sprintf(
		"(nama_barang_normalized LIKE %[1]s OR merek_normalized LIKE %[1]s OR kode_barang ILIKE %[1]s)",
		placeholder,
	))

	return where, args
}

func productLookupSelectSQL() string {
	return `
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
}

func productLookupOrderLimitSQL() string {
	return `
		ORDER BY
			nama_barang_normalized ASC,
			merek_normalized ASC,
			ukuran ASC NULLS LAST,
			id ASC
		LIMIT %s
	`
}

func scanProductLookupRows(rows pgx.Rows) ([]ports.ProductLookupItem, error) {
	items := []ports.ProductLookupItem{}

	for rows.Next() {
		item, err := scanProductLookupItem(rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

func scanProductLookupItem(row productScanner) (ports.ProductLookupItem, error) {
	var item ports.ProductLookupItem
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
		return ports.ProductLookupItem{}, err
	}

	item.Code = nullableStringPtr(code)
	item.Size = nullableIntPtr(size)

	return item, nil
}

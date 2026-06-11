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

	"github.com/jackc/pgx/v5"
)

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

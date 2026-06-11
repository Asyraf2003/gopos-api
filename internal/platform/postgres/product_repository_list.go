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
	"fmt"
	"strings"

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) List(
	ctx context.Context,
	query ports.ProductListQuery,
) ([]ports.ProductListItem, error) {
	where, args := productListWhere(query)
	page, perPage := productListPagination(query)

	args = append(args, perPage)
	limitPlaceholder := fmt.Sprintf("$%d", len(args))

	args = append(args, (page-1)*perPage)
	offsetPlaceholder := fmt.Sprintf("$%d", len(args))

	querySQL := productListSelectSQL()
	if len(where) > 0 {
		querySQL += "\n\t\tWHERE " + strings.Join(where, "\n\t\t\tAND ")
	}
	querySQL += fmt.Sprintf(
		productListOrderLimitSQL(),
		limitPlaceholder,
		offsetPlaceholder,
	)

	rows, err := r.query(ctx, querySQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanProductListRows(rows)
}

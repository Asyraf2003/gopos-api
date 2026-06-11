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
	"fmt"
	"strings"

	"pos-go/internal/modules/productcatalog/ports"
)

const (
	productListDefaultPage    = 1
	productListDefaultPerPage = 20
	productListMaxPerPage     = 100
)

func productListPagination(query ports.ProductListQuery) (int, int) {
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

	return page, perPage
}

func productListWhere(query ports.ProductListQuery) ([]string, []any) {
	where := productListStatusWhere(query.Status)
	args := []any{}

	search := strings.TrimSpace(query.Search)
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

func productListStatusWhere(status string) []string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "deleted":
		return []string{"deleted_at IS NOT NULL"}
	case "all":
		return []string{}
	default:
		return []string{"deleted_at IS NULL"}
	}
}

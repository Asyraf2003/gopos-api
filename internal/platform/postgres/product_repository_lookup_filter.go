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
	productLookupDefaultLimit = 10
	productLookupMaxLimit     = 50
)

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

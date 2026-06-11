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

import "pos-go/internal/modules/productcatalog/domain"

func productArgs(product *domain.Product) []any {
	return []any{
		product.ID(),
		nullableStringArg(product.Code()),
		product.Name(),
		product.NormalizedName(),
		product.Brand(),
		product.NormalizedBrand(),
		nullableIntArg(product.Size()),
		product.SalePriceRupiah(),
		nullableIntArg(product.ReorderPointQty()),
		nullableIntArg(product.CriticalThresholdQty()),
		nullableTimeArg(product.DeletedAt()),
		nullableNonEmptyStringArg(product.DeletedByActorID()),
		nullableNonEmptyStringArg(product.DeleteReason()),
	}
}

func nullableStringArg(value *string) any {
	if value == nil {
		return nil
	}

	return *value
}

func nullableIntArg(value *int) any {
	if value == nil {
		return nil
	}

	return *value
}

func nullableTimeArg(value any) any {
	if value == nil {
		return nil
	}

	return value
}

func nullableNonEmptyStringArg(value string) any {
	if value == "" {
		return nil
	}

	return value
}

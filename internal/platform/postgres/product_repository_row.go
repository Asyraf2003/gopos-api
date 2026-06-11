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
	"fmt"

	"pos-go/internal/modules/productcatalog/domain"
)

type productScanner interface {
	Scan(dest ...any) error
}

func productSelectSQL() string {
	return `
		SELECT
			id,
			kode_barang,
			nama_barang,
			nama_barang_normalized,
			merek,
			merek_normalized,
			ukuran,
			harga_jual,
			reorder_point_qty,
			critical_threshold_qty,
			deleted_at,
			deleted_by_actor_id,
			delete_reason
		FROM products
	`
}

func scanProduct(row productScanner) (*domain.Product, error) {
	var id string
	var code sql.NullString
	var name string
	var normalizedName string
	var brand string
	var normalizedBrand string
	var size sql.NullInt64
	var salePriceRupiah int64
	var reorderPointQty sql.NullInt64
	var criticalThresholdQty sql.NullInt64
	var deletedAt sql.NullTime
	var deletedByActorID sql.NullString
	var deleteReason sql.NullString

	err := row.Scan(
		&id,
		&code,
		&name,
		&normalizedName,
		&brand,
		&normalizedBrand,
		&size,
		&salePriceRupiah,
		&reorderPointQty,
		&criticalThresholdQty,
		&deletedAt,
		&deletedByActorID,
		&deleteReason,
	)
	if err != nil {
		return nil, err
	}

	product, err := domain.NewProduct(domain.ProductInput{
		ID:                   id,
		Code:                 nullableStringValue(code),
		Name:                 name,
		Brand:                brand,
		Size:                 nullableIntValue(size),
		SalePriceRupiah:      salePriceRupiah,
		ReorderPointQty:      nullableIntValue(reorderPointQty),
		CriticalThresholdQty: nullableIntValue(criticalThresholdQty),
	})
	if err != nil {
		return nil, err
	}

	if product.NormalizedName() != normalizedName {
		return nil, fmt.Errorf("product normalized name mismatch for id %q", id)
	}
	if product.NormalizedBrand() != normalizedBrand {
		return nil, fmt.Errorf("product normalized brand mismatch for id %q", id)
	}

	if deletedAt.Valid {
		if err := product.SoftDelete(domain.DeleteInput{
			DeletedAt:        deletedAt.Time,
			DeletedByActorID: nullableStringValue(deletedByActorID),
			Reason:           nullableStringValue(deleteReason),
		}); err != nil {
			return nil, err
		}
	}

	return product, nil
}

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

func nullableStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}

	return value.String
}

func nullableIntValue(value sql.NullInt64) *int {
	if !value.Valid {
		return nil
	}

	intValue := int(value.Int64)
	return &intValue
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

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
	"errors"
	"strings"

	"pos-go/internal/modules/supplier/domain"
	"pos-go/internal/modules/supplier/ports"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *SupplierRepository) query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return r.pool.Query(ctx, sql, args...)
}
func (r *SupplierRepository) queryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}
func (r *SupplierRepository) exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if tx, ok := TxFromContext(ctx); ok {
		return tx.Exec(ctx, sql, args...)
	}
	return r.pool.Exec(ctx, sql, args...)
}
func (r *SupplierRepository) FindByID(ctx context.Context, id domain.SupplierID) (domain.Supplier, bool, error) {
	return scanOptionalSupplier(r.queryRow(ctx, supplierSelectSQL()+" WHERE id = $1", string(id)))
}
func (r *SupplierRepository) FindByNormalizedName(ctx context.Context, name domain.NormalizedName) (domain.Supplier, bool, error) {
	sql := supplierSelectSQL() + " WHERE name_normalized = $1 ORDER BY is_active DESC, updated_at DESC, id LIMIT 1"
	return scanOptionalSupplier(r.queryRow(ctx, sql, string(name)))
}
func (r *SupplierRepository) FindActiveByNormalizedName(ctx context.Context, name domain.NormalizedName) (domain.Supplier, bool, error) {
	sql := supplierSelectSQL() + " WHERE name_normalized = $1 AND is_active = true LIMIT 1"
	return scanOptionalSupplier(r.queryRow(ctx, sql, string(name)))
}
func (r *SupplierRepository) List(ctx context.Context, filter ports.ListSuppliersFilter) ([]domain.Supplier, error) {
	query, normalizedPattern, displayPattern := supplierSearchArgs(filter.Query)
	status := map[ports.ListStatusFilter]string{ports.ListStatusInactive: "inactive", ports.ListStatusAll: "all"}[filter.Status]
	if status == "" {
		status = "active"
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	perPage := supplierBoundedLimit(filter.PerPage, 10)
	offset := (page - 1) * perPage
	sql := supplierSelectSQL() + " WHERE ($1 = '' OR name_normalized LIKE $2 OR name ILIKE $3)"
	sql += " AND ($4 = 'all' OR ($4 = 'active' AND is_active = true) OR ($4 = 'inactive' AND is_active = false))"
	sql += " ORDER BY name_normalized, id LIMIT $5 OFFSET $6"
	return r.findManySuppliers(ctx, sql, query, normalizedPattern, displayPattern, status, perPage, offset)
}
func (r *SupplierRepository) Lookup(ctx context.Context, filter ports.LookupSuppliersFilter) ([]domain.Supplier, error) {
	query, normalizedPattern, displayPattern := supplierSearchArgs(filter.Query)
	limit := supplierBoundedLimit(filter.Limit, 20)
	sql := supplierSelectSQL() + " WHERE ($1 = '' OR name_normalized LIKE $2 OR name ILIKE $3)"
	sql += " AND ($4 = false OR is_active = true) ORDER BY name_normalized, id LIMIT $5"
	return r.findManySuppliers(ctx, sql, query, normalizedPattern, displayPattern, filter.ActiveOnly, limit)
}
func scanOptionalSupplier(row supplierScanner) (domain.Supplier, bool, error) {
	supplier, err := scanSupplier(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Supplier{}, false, nil
	}
	if err != nil {
		return domain.Supplier{}, false, err
	}
	return supplier, true, nil
}
func supplierSearchArgs(query string) (string, string, string) {
	query = strings.TrimSpace(query)
	return query, "%" + string(domain.NormalizeName(query)) + "%", "%" + query + "%"
}
func supplierBoundedLimit(value int, fallback int) int {
	if value <= 0 {
		return fallback
	}
	if value > 50 {
		return 50
	}
	return value
}
func (r *SupplierRepository) findManySuppliers(ctx context.Context, sql string, args ...any) ([]domain.Supplier, error) {
	rows, err := r.query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Supplier, error) {
		return scanSupplier(row)
	})
}

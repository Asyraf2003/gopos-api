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

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"

	"github.com/jackc/pgx/v5"
)

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	row := r.queryRow(ctx, productSelectSQL()+`
		WHERE id = $1
	`, id)

	product, err := scanProduct(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ports.ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	_ = ctx
	_ = id

	return nil, errProductRepositoryNotImplemented
}

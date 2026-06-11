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

	"pos-go/internal/modules/productcatalog/ports"
)

func (r *ProductRepository) CheckCreateDuplicate(
	ctx context.Context,
	candidate ports.ProductDuplicateCandidate,
) error {
	if err := r.checkProductDuplicateCode(ctx, "", candidate); err != nil {
		return err
	}

	return r.checkProductDuplicateIdentity(ctx, "", candidate)
}

func (r *ProductRepository) CheckUpdateDuplicate(
	ctx context.Context,
	productID string,
	candidate ports.ProductDuplicateCandidate,
) error {
	if err := r.checkProductDuplicateCode(ctx, productID, candidate); err != nil {
		return err
	}

	return r.checkProductDuplicateIdentity(ctx, productID, candidate)
}

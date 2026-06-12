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

package http

import (
	"context"

	productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"
)

type listProducts interface {
	Execute(context.Context, productcatalogusecase.ListProductsQuery) (productcatalogusecase.ListProductsResult, error)
}

type lookupProducts interface {
	Execute(context.Context, productcatalogusecase.LookupProductsQuery) (productcatalogusecase.LookupProductsResult, error)
}

type getProductDetail interface {
	Execute(context.Context, productcatalogusecase.GetProductDetailQuery) (productcatalogusecase.GetProductDetailResult, error)
}

type createProduct interface {
	Execute(context.Context, productcatalogusecase.CreateProductCommand) (productcatalogusecase.CreateProductResult, error)
}

type updateProduct interface {
	Execute(context.Context, productcatalogusecase.UpdateProductCommand) (productcatalogusecase.UpdateProductResult, error)
}

type softDeleteProduct interface {
	Execute(context.Context, productcatalogusecase.SoftDeleteProductCommand) (productcatalogusecase.SoftDeleteProductResult, error)
}

type restoreProduct interface {
	Execute(context.Context, productcatalogusecase.RestoreProductCommand) (productcatalogusecase.RestoreProductResult, error)
}

type listProductVersions interface {
	Execute(context.Context, productcatalogusecase.ListProductVersionsQuery) (productcatalogusecase.ListProductVersionsResult, error)
}

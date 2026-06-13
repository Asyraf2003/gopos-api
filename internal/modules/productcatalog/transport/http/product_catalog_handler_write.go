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
	stdhttp "net/http"

	productcatalogusecase "pos-go/internal/modules/productcatalog/usecase"
	productcatalogid "pos-go/internal/presentation/http/id/productcatalog"
	httpresponse "pos-go/internal/transport/http/response"

	"github.com/labstack/echo/v4"
)

func (h ProductCatalogHandler) Create(c echo.Context) error {
	var req productUpsertRequest
	if err := c.Bind(&req); err != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_request_body", "invalid request body")
	}

	result, err := h.create.Execute(c.Request().Context(), productcatalogusecase.CreateProductCommand{
		Code:                 req.Code,
		Name:                 req.Name,
		Brand:                req.Brand,
		Size:                 req.Size,
		SalePriceRupiah:      req.SalePriceRupiah,
		ReorderPointQty:      req.ReorderPointQty,
		CriticalThresholdQty: req.CriticalThresholdQty,
		ActorID:              actorIDFromRequest(c),
		Reason:               req.Reason,
	})
	if err != nil {
		return mapProductCatalogError(err)
	}

	return c.JSON(stdhttp.StatusCreated, httpresponse.Success(productcatalogid.FromCreatedProduct(result)))
}

func (h ProductCatalogHandler) Update(c echo.Context) error {
	var req productUpsertRequest
	if err := c.Bind(&req); err != nil {
		return httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_request_body", "invalid request body")
	}

	result, err := h.update.Execute(c.Request().Context(), productcatalogusecase.UpdateProductCommand{
		ID:                   c.Param("id"),
		Code:                 req.Code,
		Name:                 req.Name,
		Brand:                req.Brand,
		Size:                 req.Size,
		SalePriceRupiah:      req.SalePriceRupiah,
		ReorderPointQty:      req.ReorderPointQty,
		CriticalThresholdQty: req.CriticalThresholdQty,
		ActorID:              actorIDFromRequest(c),
		Reason:               req.Reason,
	})
	if err != nil {
		return mapProductCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(productcatalogid.FromUpdatedProduct(result)))
}

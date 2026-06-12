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

	"github.com/labstack/echo/v4"
)

func (h ProductCatalogHandler) Create(c echo.Context) error {
	var req productUpsertRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
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

	return c.JSON(stdhttp.StatusCreated, successEnvelope(productcatalogid.FromCreatedProduct(result)))
}

func (h ProductCatalogHandler) Update(c echo.Context) error {
	var req productUpsertRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
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

	return c.JSON(stdhttp.StatusOK, successEnvelope(productcatalogid.FromUpdatedProduct(result)))
}

func (h ProductCatalogHandler) Delete(c echo.Context) error {
	req, err := bindProductLifecycleRequest(c)
	if err != nil {
		return err
	}

	result, err := h.softDelete.Execute(c.Request().Context(), productcatalogusecase.SoftDeleteProductCommand{
		ID:      c.Param("id"),
		ActorID: actorIDFromRequest(c),
		Reason:  req.Reason,
	})
	if err != nil {
		return mapProductCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(productcatalogid.FromDeletedProduct(result)))
}

func (h ProductCatalogHandler) Restore(c echo.Context) error {
	req, err := bindProductLifecycleRequest(c)
	if err != nil {
		return err
	}

	result, err := h.restore.Execute(c.Request().Context(), productcatalogusecase.RestoreProductCommand{
		ID:      c.Param("id"),
		ActorID: actorIDFromRequest(c),
		Reason:  req.Reason,
	})
	if err != nil {
		return mapProductCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(productcatalogid.FromRestoredProduct(result)))
}

func bindProductLifecycleRequest(c echo.Context) (productLifecycleRequest, error) {
	var req productLifecycleRequest
	if c.Request().ContentLength == 0 {
		return req, nil
	}

	if err := c.Bind(&req); err != nil {
		return productLifecycleRequest{}, echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
	}

	return req, nil
}

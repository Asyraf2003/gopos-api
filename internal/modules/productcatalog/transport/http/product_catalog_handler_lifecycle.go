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

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(productcatalogid.FromDeletedProduct(result)))
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

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(productcatalogid.FromRestoredProduct(result)))
}

func bindProductLifecycleRequest(c echo.Context) (productLifecycleRequest, error) {
	var req productLifecycleRequest
	if c.Request().ContentLength == 0 {
		return req, nil
	}

	if err := c.Bind(&req); err != nil {
		return productLifecycleRequest{}, httpresponse.NewHTTPError(stdhttp.StatusBadRequest, "invalid_request_body", "invalid request body")
	}

	return req, nil
}

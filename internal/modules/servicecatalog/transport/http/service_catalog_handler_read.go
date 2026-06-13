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
	httpresponse "pos-go/internal/transport/http/response"

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
	servicecatalogid "pos-go/internal/presentation/http/id/servicecatalog"

	"github.com/labstack/echo/v4"
)

func (h ServiceCatalogHandler) List(c echo.Context) error {
	status, err := parseListStatus(c.QueryParam("status"))
	if err != nil {
		return err
	}

	page, err := parseOptionalIntQuery(c, "page")
	if err != nil {
		return err
	}

	perPage, err := parseOptionalIntQuery(c, "per_page")
	if err != nil {
		return err
	}

	results, err := h.list.Execute(c.Request().Context(), servicecatalogusecase.ListServiceCatalogItemsCommand{
		Query:   c.QueryParam("q"),
		Status:  status,
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(servicecatalogid.FromServiceCatalogItems(results)))
}

func (h ServiceCatalogHandler) Lookup(c echo.Context) error {
	limit, err := parseOptionalIntQuery(c, "limit")
	if err != nil {
		return err
	}

	includeInactive, err := parseLookupIncludeInactive(c)
	if err != nil {
		return err
	}

	results, err := h.lookup.Execute(c.Request().Context(), servicecatalogusecase.LookupServiceCatalogItemsCommand{
		Query:           c.QueryParam("q"),
		Limit:           limit,
		IncludeInactive: includeInactive,
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(servicecatalogid.FromServiceCatalogLookups(results)))
}

func (h ServiceCatalogHandler) Show(c echo.Context) error {
	result, err := h.show.Execute(c.Request().Context(), servicecatalogusecase.ShowServiceCatalogItemCommand{
		ID: c.Param("id"),
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, httpresponse.Success(servicecatalogid.FromServiceCatalogItem(result)))
}

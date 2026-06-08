package http

import (
	stdhttp "net/http"

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

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogItems(results)))
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

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogLookups(results)))
}

func (h ServiceCatalogHandler) Show(c echo.Context) error {
	result, err := h.show.Execute(c.Request().Context(), servicecatalogusecase.ShowServiceCatalogItemCommand{
		ID: c.Param("id"),
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogItem(result)))
}

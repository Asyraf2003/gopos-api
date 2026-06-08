package http

import (
	stdhttp "net/http"

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"
	servicecatalogid "pos-go/internal/presentation/http/id/servicecatalog"

	"github.com/labstack/echo/v4"
)

func (h ServiceCatalogHandler) Create(c echo.Context) error {
	var req serviceCatalogUpsertRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
	}

	result, err := h.create.Execute(c.Request().Context(), servicecatalogusecase.CreateServiceCatalogItemCommand{
		Name:               req.Name,
		DefaultPriceRupiah: req.DefaultPriceRupiah,
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusCreated, successEnvelope(servicecatalogid.FromServiceCatalogItem(result)))
}

func (h ServiceCatalogHandler) Update(c echo.Context) error {
	var req serviceCatalogUpsertRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
	}

	result, err := h.update.Execute(c.Request().Context(), servicecatalogusecase.UpdateServiceCatalogItemCommand{
		ID:                 c.Param("id"),
		Name:               req.Name,
		DefaultPriceRupiah: req.DefaultPriceRupiah,
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogItem(result)))
}

func (h ServiceCatalogHandler) Activate(c echo.Context) error {
	result, err := h.activate.Execute(c.Request().Context(), servicecatalogusecase.ActivateServiceCatalogItemCommand{
		ID: c.Param("id"),
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogItem(result)))
}

func (h ServiceCatalogHandler) Deactivate(c echo.Context) error {
	var req serviceCatalogDeactivateRequest
	if c.Request().ContentLength != 0 {
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(stdhttp.StatusBadRequest, "invalid request body")
		}
	}
	_ = req.Reason

	result, err := h.deactivate.Execute(c.Request().Context(), servicecatalogusecase.DeactivateServiceCatalogItemCommand{
		ID: c.Param("id"),
	})
	if err != nil {
		return mapServiceCatalogError(err)
	}

	return c.JSON(stdhttp.StatusOK, successEnvelope(servicecatalogid.FromServiceCatalogItem(result)))
}

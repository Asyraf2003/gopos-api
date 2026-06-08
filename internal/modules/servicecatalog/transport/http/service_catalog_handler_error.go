package http

import (
	"errors"

	"pos-go/internal/modules/servicecatalog/domain"
	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"

	"github.com/labstack/echo/v4"
)

func mapServiceCatalogError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, servicecatalogusecase.ErrServiceCatalogItemNotFound):
		return echo.NewHTTPError(404, "service catalog item not found")
	case errors.Is(err, servicecatalogusecase.ErrDuplicateServiceCatalogItemNormalizedName):
		return echo.NewHTTPError(409, "service catalog item name already exists")
	case errors.Is(err, servicecatalogusecase.ErrInvalidLookupLimit):
		return echo.NewHTTPError(400, "lookup limit must be between 1 and 50")
	case errors.Is(err, domain.ErrInvalidServiceCatalogItemID),
		errors.Is(err, domain.ErrBlankServiceCatalogItemName),
		errors.Is(err, domain.ErrInvalidServiceCatalogItemDefaultPrice):
		return echo.NewHTTPError(400, err.Error())
	default:
		return echo.NewHTTPError(500, "service catalog request failed")
	}
}

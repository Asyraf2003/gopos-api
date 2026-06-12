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
	"errors"

	"pos-go/internal/modules/productcatalog/domain"
	"pos-go/internal/modules/productcatalog/ports"

	"github.com/labstack/echo/v4"
)

func mapProductCatalogError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ports.ErrProductNotFound):
		return echo.NewHTTPError(404, "product not found")
	case errors.Is(err, ports.ErrDuplicateProductCode):
		return echo.NewHTTPError(409, "product code already exists")
	case errors.Is(err, ports.ErrDuplicateProductIdentity):
		return echo.NewHTTPError(409, "product identity already exists")
	case errors.Is(err, domain.ErrProductIDRequired),
		errors.Is(err, domain.ErrProductNameRequired),
		errors.Is(err, domain.ErrProductBrandRequired),
		errors.Is(err, domain.ErrProductSalePriceMustBePositive),
		errors.Is(err, domain.ErrProductThresholdPairRequired),
		errors.Is(err, domain.ErrProductThresholdNegative),
		errors.Is(err, domain.ErrProductCriticalAboveReorder),
		errors.Is(err, domain.ErrProductDeleteTimeRequired),
		errors.Is(err, domain.ErrProductAlreadyDeleted),
		errors.Is(err, domain.ErrProductNotDeleted):
		return echo.NewHTTPError(400, err.Error())
	default:
		return echo.NewHTTPError(500, "product catalog request failed")
	}
}

package http

import (
	"strconv"
	"strings"

	"pos-go/internal/modules/servicecatalog/ports"

	"github.com/labstack/echo/v4"
)

type serviceCatalogUpsertRequest struct {
	Name               string `json:"name"`
	DefaultPriceRupiah int64  `json:"default_price_rupiah"`
}

type serviceCatalogDeactivateRequest struct {
	Reason string `json:"reason"`
}

func parseOptionalIntQuery(c echo.Context, name string) (int, error) {
	raw := strings.TrimSpace(c.QueryParam(name))
	if raw == "" {
		return 0, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, echo.NewHTTPError(400, name+" must be an integer")
	}

	return value, nil
}

func parseListStatus(raw string) (ports.ListStatusFilter, error) {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "":
		return "", nil
	case string(ports.ListStatusActive):
		return ports.ListStatusActive, nil
	case string(ports.ListStatusInactive):
		return ports.ListStatusInactive, nil
	case string(ports.ListStatusAll):
		return ports.ListStatusAll, nil
	default:
		return "", echo.NewHTTPError(400, "status must be active, inactive, or all")
	}
}

func parseLookupIncludeInactive(c echo.Context) (bool, error) {
	raw := strings.TrimSpace(c.QueryParam("active_only"))
	if raw == "" {
		return false, nil
	}

	activeOnly, err := strconv.ParseBool(raw)
	if err != nil {
		return false, echo.NewHTTPError(400, "active_only must be a boolean")
	}

	return !activeOnly, nil
}

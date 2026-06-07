package http

import (
	"errors"
	"net/http"

	"pos-go/internal/modules/capability/ports"

	"github.com/labstack/echo/v4"
)

type responseEnvelope struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type disableCapabilityRequest struct {
	Reason string `json:"reason"`
}

func capabilityHTTPError(err error) error {
	if errors.Is(err, ports.ErrCapabilityNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "capability not found")
	}

	return err
}

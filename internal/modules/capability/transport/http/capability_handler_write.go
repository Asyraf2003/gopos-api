package http

import (
	"net/http"
	"strings"

	capabilitypresenter "pos-go/internal/presentation/http/id/capability"

	"github.com/labstack/echo/v4"
)

func (h *CapabilityHandler) Enable(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	if err := h.enableUsecase.Execute(c.Request().Context(), key); err != nil {
		return capabilityHTTPError(err)
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomain(capability),
	})
}

func (h *CapabilityHandler) Disable(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
	}

	var req disableCapabilityRequest
	if c.Request().Body != nil && c.Request().ContentLength != 0 {
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
		}
	}

	if err := h.disableUsecase.Execute(c.Request().Context(), key, req.Reason); err != nil {
		return capabilityHTTPError(err)
	}

	capability, err := h.showUsecase.Execute(c.Request().Context(), key)
	if err != nil {
		return capabilityHTTPError(err)
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomain(capability),
	})
}

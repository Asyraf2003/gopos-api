package http

import (
	"net/http"
	"strings"

	capabilitypresenter "pos-go/internal/presentation/http/id/capability"

	"github.com/labstack/echo/v4"
)

func (h *CapabilityHandler) List(c echo.Context) error {
	capabilities, err := h.listUsecase.Execute(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseEnvelope{
		Success: true,
		Data:    capabilitypresenter.FromDomainList(capabilities),
	})
}

func (h *CapabilityHandler) Show(c echo.Context) error {
	key := strings.TrimSpace(c.Param("key"))
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "capability key is required")
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

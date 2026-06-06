package http

import (
	"net/http"

	"pos-go/internal/modules/system/ports"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	checker ports.HealthChecker
}

func NewHealthHandler(checker ports.HealthChecker) HealthHandler {
	return HealthHandler{checker: checker}
}

func (h HealthHandler) Register(group *echo.Group) {
	group.GET("/health", h.Get)
}

func (h HealthHandler) Get(c echo.Context) error {
	if err := h.checker.Check(c.Request().Context()); err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]any{
			"status":   "degraded",
			"database": "down",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":   "ok",
		"database": "up",
	})
}

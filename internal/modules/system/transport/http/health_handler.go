package http

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) HealthHandler {
	return HealthHandler{pool: pool}
}

func (h HealthHandler) Register(group *echo.Group) {
	group.GET("/health", h.Get)
}

func (h HealthHandler) Get(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	if err := h.pool.Ping(ctx); err != nil {
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

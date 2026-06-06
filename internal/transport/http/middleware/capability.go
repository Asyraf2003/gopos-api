package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"pos-go/internal/modules/capability/domain"

	"github.com/labstack/echo/v4"
)

type CapabilityChecker interface {
	Execute(ctx context.Context, key string) error
}

type capabilityCheckerFunc func(ctx context.Context, key string) error

func (f capabilityCheckerFunc) Execute(ctx context.Context, key string) error {
	return f(ctx, key)
}

func RequireCapability(capabilityKey string, checker CapabilityChecker) echo.MiddlewareFunc {
	capabilityKey = strings.TrimSpace(capabilityKey)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if capabilityKey == "" || checker == nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "capability guard misconfigured")
			}

			if err := checker.Execute(c.Request().Context(), capabilityKey); err != nil {
				if errors.Is(err, domain.ErrCapabilityDisabled) {
					return echo.NewHTTPError(http.StatusForbidden, "capability disabled")
				}

				return echo.NewHTTPError(http.StatusInternalServerError, "capability check failed")
			}

			return next(c)
		}
	}
}

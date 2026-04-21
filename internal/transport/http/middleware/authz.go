package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RequirePermission(permissionKey string) echo.MiddlewareFunc {
	permissionKey = strings.TrimSpace(permissionKey)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if permissionKey == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, "permission guard misconfigured")
			}

			principal, ok := PrincipalFromContext(c.Request().Context())
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "authentication required")
			}

			if !principal.HasPermission(permissionKey) {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}

			return next(c)
		}
	}
}

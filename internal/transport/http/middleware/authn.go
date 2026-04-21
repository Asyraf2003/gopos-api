package middleware

import (
	"context"
	"net/http"
	"strings"

	"pos-go/internal/modules/auth/domain"
	"pos-go/internal/modules/auth/ports"

	"github.com/labstack/echo/v4"
)

type principalContextKey struct{}

func WithPrincipal(ctx context.Context, principal domain.Principal) context.Context {
	return context.WithValue(ctx, principalContextKey{}, principal)
}

func PrincipalFromContext(ctx context.Context) (domain.Principal, bool) {
	principal, ok := ctx.Value(principalContextKey{}).(domain.Principal)
	return principal, ok
}

func RequireAuth(
	verifier ports.AccessTokenVerifier,
	resolver ports.PrincipalResolver,
	sessionChecker ports.SessionStatusChecker,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authz := strings.TrimSpace(c.Request().Header.Get(echo.HeaderAuthorization))
			if authz == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing bearer token")
			}

			const prefix = "Bearer "
			if !strings.HasPrefix(authz, prefix) {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid bearer token")
			}

			token := strings.TrimSpace(strings.TrimPrefix(authz, prefix))
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid bearer token")
			}

			claims, err := verifier.VerifyAccessToken(c.Request().Context(), token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid access token")
			}

			if sessionChecker != nil {
				active, err := sessionChecker.IsSessionActive(c.Request().Context(), claims.SessionID)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "session check failed")
				}
				if !active {
					return echo.NewHTTPError(http.StatusUnauthorized, "inactive session")
				}
			}

			principal, err := resolver.Resolve(c.Request().Context(), ports.ResolvePrincipalInput{
				AccountID:  claims.AccountID,
				SessionID:  claims.SessionID,
				TrustLevel: claims.TrustLevel,
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "principal resolve failed")
			}

			req := c.Request().WithContext(WithPrincipal(c.Request().Context(), principal))
			c.SetRequest(req)

			return next(c)
		}
	}
}

package http

import (
	"net/http"

	authdomain "pos-go/internal/modules/auth/domain"
	authusecase "pos-go/internal/modules/auth/usecase"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

type LogoutHandler struct {
	usecase *authusecase.LogoutCurrentSession
}

func NewLogoutHandler(usecase *authusecase.LogoutCurrentSession) *LogoutHandler {
	return &LogoutHandler{usecase: usecase}
}

func (h *LogoutHandler) Register(group *echo.Group) {
	group.POST("/logout", h.Logout)
}

func (h *LogoutHandler) Logout(c echo.Context) error {
	principal, ok := httpmw.PrincipalFromContext(c.Request().Context())
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "authentication required")
	}

	if err := h.usecase.Execute(c.Request().Context(), authdomain.Principal{
		AccountID:   principal.AccountID,
		SessionID:   principal.SessionID,
		Roles:       principal.Roles,
		Permissions: principal.Permissions,
		TrustLevel:  principal.TrustLevel,
	}); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

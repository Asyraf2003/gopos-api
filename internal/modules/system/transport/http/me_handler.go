package http

import (
	"net/http"

	systempresenter "pos-go/internal/presentation/http/id/system"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
)

type MeHandler struct{}

func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

func (h *MeHandler) Register(group *echo.Group) {
	group.GET("/me", h.Show)
}

func (h *MeHandler) Show(c echo.Context) error {
	principal, ok := httpmw.PrincipalFromContext(c.Request().Context())
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "authentication required")
	}

	return c.JSON(http.StatusOK, systempresenter.Me(principal))
}

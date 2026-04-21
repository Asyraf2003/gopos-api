package http

import (
	"context"
	"net/http"
	"strings"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type AssignAccountRoleUsecase interface {
	Execute(ctx context.Context, accountID string, roleKey string) error
}

type RemoveAccountRoleUsecase interface {
	Execute(ctx context.Context, accountID string, roleKey string) error
}

type AccountRoleHandler struct {
	assignUsecase AssignAccountRoleUsecase
	removeUsecase RemoveAccountRoleUsecase
}

func NewAccountRoleHandler(
	assignUsecase AssignAccountRoleUsecase,
	removeUsecase RemoveAccountRoleUsecase,
) *AccountRoleHandler {
	return &AccountRoleHandler{
		assignUsecase: assignUsecase,
		removeUsecase: removeUsecase,
	}
}

func (h *AccountRoleHandler) Register(group *echo.Group) {
	group.POST("/accounts/:account_id/roles", h.Assign)
	group.DELETE("/accounts/:account_id/roles/:role_key", h.Remove)
}

type assignAccountRoleRequest struct {
	RoleKey string `json:"role_key"`
}

func (h *AccountRoleHandler) Assign(c echo.Context) error {
	var req assignAccountRoleRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	accountID := strings.TrimSpace(c.Param("account_id"))
	roleKey := strings.TrimSpace(req.RoleKey)

	if accountID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "account id is required")
	}
	if roleKey == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "role key is required")
	}

	if err := h.assignUsecase.Execute(c.Request().Context(), accountID, roleKey); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *AccountRoleHandler) Remove(c echo.Context) error {
	accountID := strings.TrimSpace(c.Param("account_id"))
	roleKey := strings.TrimSpace(c.Param("role_key"))

	if accountID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "account id is required")
	}
	if roleKey == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "role key is required")
	}

	if err := h.removeUsecase.Execute(c.Request().Context(), accountID, roleKey); err != nil {
		if err == authusecase.ErrBaseRoleRemovalNotAllowed {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

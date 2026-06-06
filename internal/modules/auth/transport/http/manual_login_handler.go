package http

import (
	"context"
	"net/http"

	authusecase "pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type ManualLoginUsecase interface {
	Execute(ctx context.Context, in authusecase.ManualLoginInput) (authusecase.ManualLoginOutput, error)
}

type ManualLoginHandler struct {
	usecase ManualLoginUsecase
}

func NewManualLoginHandler(usecase ManualLoginUsecase) *ManualLoginHandler {
	return &ManualLoginHandler{usecase: usecase}
}

func (h *ManualLoginHandler) Register(group *echo.Group) {
	group.POST("/manual/login", h.Login)
}

type manualLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *ManualLoginHandler) Login(c echo.Context) error {
	var req manualLoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	out, err := h.usecase.Execute(c.Request().Context(), authusecase.ManualLoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if err == authusecase.ErrManualLoginInvalidCredentials {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid manual login credentials")
		}
		return err
	}

	return c.JSON(http.StatusOK, out)
}

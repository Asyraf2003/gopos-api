package http

import (
	"log"
	"net/http"
	"strings"

	"pos-go/internal/modules/auth/usecase"

	"github.com/labstack/echo/v4"
)

type GoogleFlow interface {
	GoogleStart(ctx echo.Context, in usecase.GoogleStartInput) (usecase.GoogleStartOutput, error)
	GoogleCallback(ctx echo.Context, in usecase.GoogleCallbackInput) (usecase.GoogleCallbackOutput, error)
}

type googleFlowAdapter struct {
	flow *usecase.GoogleFlow
}

func NewGoogleFlowAdapter(flow *usecase.GoogleFlow) GoogleFlow {
	return &googleFlowAdapter{flow: flow}
}

func (a *googleFlowAdapter) GoogleStart(c echo.Context, in usecase.GoogleStartInput) (usecase.GoogleStartOutput, error) {
	return a.flow.GoogleStart(c.Request().Context(), in)
}

func (a *googleFlowAdapter) GoogleCallback(c echo.Context, in usecase.GoogleCallbackInput) (usecase.GoogleCallbackOutput, error) {
	return a.flow.GoogleCallback(c.Request().Context(), in)
}

type GoogleHandler struct {
	flow        GoogleFlow
	redirectURL string
}

func NewGoogleHandler(flow GoogleFlow, redirectURL string) *GoogleHandler {
	return &GoogleHandler{
		flow:        flow,
		redirectURL: strings.TrimSpace(redirectURL),
	}
}

func (h *GoogleHandler) Register(group *echo.Group) {
	group.GET("/google/start", h.Start)
	group.GET("/google/callback", h.Callback)
}

func (h *GoogleHandler) Start(c echo.Context) error {
	redirectURL := strings.TrimSpace(c.QueryParam("redirect_url"))
	if redirectURL == "" {
		redirectURL = h.redirectURL
	}

	out, err := h.flow.GoogleStart(c, usecase.GoogleStartInput{
		Purpose:     c.QueryParam("purpose"),
		RedirectURL: redirectURL,
	})
	if err != nil {
		log.Printf("auth google start error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *GoogleHandler) Callback(c echo.Context) error {
	out, err := h.flow.GoogleCallback(c, usecase.GoogleCallbackInput{
		Code:        c.QueryParam("code"),
		State:       c.QueryParam("state"),
		RedirectURL: h.redirectURL,
		Client: usecase.ClientInfo{
			UserAgent: c.Request().UserAgent(),
			IP:        c.RealIP(),
		},
	})
	if err != nil {
		log.Printf("auth google callback error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, out)
}

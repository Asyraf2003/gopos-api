package http

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCapabilityHandler_ShowRejectsMissingKey(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodGet, "/api/admin/capabilities/", "")
	c.SetParamNames("key")
	c.SetParamValues(" ")
	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusBadRequest)
	assertEqual(t, fake.showCalls, 0, "show calls")
}

func TestCapabilityHandler_ShowMapsNotFound(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodGet, "/api/admin/capabilities/missing", "")
	c.SetParamNames("key")
	c.SetParamValues("missing")
	handler, _ := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusNotFound)
}

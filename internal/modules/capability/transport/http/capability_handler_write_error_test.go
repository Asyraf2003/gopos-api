package http

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCapabilityHandler_DisableRejectsInvalidBody(t *testing.T) {
	e := echo.New()
	c, _ := newCapabilityContext(e, http.MethodPost, "/api/admin/capabilities/capability.manage/disable", `{"reason":`)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")
	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Disable(c)
	if err == nil {
		t.Fatal("Disable() error = nil, want error")
	}
	assertHTTPError(t, err, http.StatusBadRequest)
	assertEqual(t, fake.disableCalls, 0, "disable calls")
}

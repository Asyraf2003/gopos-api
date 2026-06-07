package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCapabilityHandler_ListSuccess(t *testing.T) {
	e := echo.New()
	c, rec := newCapabilityContext(e, http.MethodGet, "/api/admin/capabilities", "")
	handler, fake := newCapabilityHandlerForTest(t)

	if err := handler.List(c); err != nil {
		t.Fatalf("List() error = %v", err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertEqual(t, fake.listCalls, 1, "list calls")
	envelope := decodeEnvelope(t, rec)
	if !envelope.Success {
		t.Fatal("success = false, want true")
	}
	data := decodeCapabilityList(t, envelope.Data)
	assertEqual(t, len(data), 1, "data len")
	assertEqual(t, data[0].Key, "capability.manage", "key")
}

func TestCapabilityHandler_ShowSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/admin/capabilities/capability.manage", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")
	handler, fake := newCapabilityHandlerForTest(t)

	if err := handler.Show(c); err != nil {
		t.Fatalf("Show() error = %v", err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertEqual(t, fake.showCalls, 1, "show calls")
	data := decodeCapability(t, decodeEnvelope(t, rec).Data)
	assertEqual(t, data.Key, "capability.manage", "key")
	assertEqual(t, data.RiskLevel, "high", "risk_level")
}

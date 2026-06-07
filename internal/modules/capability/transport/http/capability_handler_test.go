package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"pos-go/internal/modules/capability/domain"
	"pos-go/internal/modules/capability/ports"

	"github.com/labstack/echo/v4"
)

func TestCapabilityHandler_ListSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/admin/capabilities", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler, fake := newCapabilityHandlerForTest(t)

	if err := handler.List(c); err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if fake.listCalls != 1 {
		t.Fatalf("list calls = %d, want 1", fake.listCalls)
	}

	var envelope testEnvelope
	decodeResponse(t, rec.Body.String(), &envelope)
	if !envelope.Success {
		t.Fatal("success = false, want true")
	}

	var data []capabilityResponseForTest
	decodeRawData(t, envelope.Data, &data)
	if len(data) != 1 {
		t.Fatalf("data len = %d, want 1", len(data))
	}
	if data[0].Key != "capability.manage" {
		t.Fatalf("key = %q, want capability.manage", data[0].Key)
	}
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

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if fake.showCalls != 1 {
		t.Fatalf("show calls = %d, want 1", fake.showCalls)
	}

	var envelope testEnvelope
	decodeResponse(t, rec.Body.String(), &envelope)
	var data capabilityResponseForTest
	decodeRawData(t, envelope.Data, &data)
	if data.Key != "capability.manage" {
		t.Fatalf("key = %q, want capability.manage", data.Key)
	}
	if data.RiskLevel != "high" {
		t.Fatalf("risk_level = %q, want high", data.RiskLevel)
	}
}

func TestCapabilityHandler_ShowRejectsMissingKey(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/admin/capabilities/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues(" ")

	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
	if fake.showCalls != 0 {
		t.Fatalf("show calls = %d, want 0", fake.showCalls)
	}
}

func TestCapabilityHandler_ShowMapsNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/admin/capabilities/missing", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("missing")

	handler, _ := newCapabilityHandlerForTest(t)

	err := handler.Show(c)
	if err == nil {
		t.Fatal("Show() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", httpErr.Code)
	}
}

func TestCapabilityHandler_EnableSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/capabilities/capability.manage/enable", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")

	handler, fake := newCapabilityHandlerForTest(t)
	fake.capabilities["capability.manage"] = fake.capabilities["capability.manage"].Disable("maintenance")

	if err := handler.Enable(c); err != nil {
		t.Fatalf("Enable() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if fake.enableCalls != 1 {
		t.Fatalf("enable calls = %d, want 1", fake.enableCalls)
	}

	var envelope testEnvelope
	decodeResponse(t, rec.Body.String(), &envelope)
	var data capabilityResponseForTest
	decodeRawData(t, envelope.Data, &data)
	if !data.Enabled {
		t.Fatal("enabled = false, want true")
	}
	if data.DisabledReason != "" {
		t.Fatalf("disabled_reason = %q, want empty", data.DisabledReason)
	}
}

func TestCapabilityHandler_DisableSuccessWithReason(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/capabilities/capability.manage/disable", strings.NewReader(`{"reason":"maintenance"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")

	handler, fake := newCapabilityHandlerForTest(t)

	if err := handler.Disable(c); err != nil {
		t.Fatalf("Disable() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if fake.disableCalls != 1 {
		t.Fatalf("disable calls = %d, want 1", fake.disableCalls)
	}
	if fake.lastDisableReason != "maintenance" {
		t.Fatalf("reason = %q, want maintenance", fake.lastDisableReason)
	}

	var envelope testEnvelope
	decodeResponse(t, rec.Body.String(), &envelope)
	var data capabilityResponseForTest
	decodeRawData(t, envelope.Data, &data)
	if data.Enabled {
		t.Fatal("enabled = true, want false")
	}
	if data.DisabledReason != "maintenance" {
		t.Fatalf("disabled_reason = %q, want maintenance", data.DisabledReason)
	}
}

func TestCapabilityHandler_DisableRejectsInvalidBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/admin/capabilities/capability.manage/disable", strings.NewReader(`{"reason":`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("key")
	c.SetParamValues("capability.manage")

	handler, fake := newCapabilityHandlerForTest(t)

	err := handler.Disable(c)
	if err == nil {
		t.Fatal("Disable() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", httpErr.Code)
	}
	if fake.disableCalls != 0 {
		t.Fatalf("disable calls = %d, want 0", fake.disableCalls)
	}
}

func TestCapabilityHandler_DoesNotImportPostgresAdapter(t *testing.T) {
	source, err := os.ReadFile("capability_handler.go")
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if strings.Contains(string(source), "internal/platform/postgres") {
		t.Fatal("handler imports PostgreSQL adapter")
	}
}

type fakeCapabilityUsecases struct {
	capabilities      map[string]domain.Capability
	listCalls         int
	showCalls         int
	enableCalls       int
	disableCalls      int
	lastDisableReason string
	err               error
}

func (f *fakeCapabilityUsecases) Execute(ctx context.Context) ([]domain.Capability, error) {
	_ = ctx
	f.listCalls++
	if f.err != nil {
		return nil, f.err
	}

	out := make([]domain.Capability, 0, len(f.capabilities))
	for _, capability := range f.capabilities {
		out = append(out, capability)
	}

	return out, nil
}

func (f *fakeCapabilityUsecases) ExecuteShow(ctx context.Context, key string) (domain.Capability, error) {
	_ = ctx
	f.showCalls++
	if f.err != nil {
		return domain.Capability{}, f.err
	}

	capability, ok := f.capabilities[key]
	if !ok {
		return domain.Capability{}, ports.ErrCapabilityNotFound
	}

	return capability, nil
}

func (f *fakeCapabilityUsecases) ExecuteEnable(ctx context.Context, key string) error {
	_ = ctx
	f.enableCalls++
	if f.err != nil {
		return f.err
	}

	capability, ok := f.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.capabilities[key] = capability.Enable()

	return nil
}

func (f *fakeCapabilityUsecases) ExecuteDisable(ctx context.Context, key string, reason string) error {
	_ = ctx
	f.disableCalls++
	f.lastDisableReason = reason
	if f.err != nil {
		return f.err
	}

	capability, ok := f.capabilities[key]
	if !ok {
		return ports.ErrCapabilityNotFound
	}
	f.capabilities[key] = capability.Disable(reason)

	return nil
}

type fakeShowCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeShowCapabilityUsecase) Execute(ctx context.Context, key string) (domain.Capability, error) {
	return f.fake.ExecuteShow(ctx, key)
}

type fakeEnableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeEnableCapabilityUsecase) Execute(ctx context.Context, key string) error {
	return f.fake.ExecuteEnable(ctx, key)
}

type fakeDisableCapabilityUsecase struct {
	fake *fakeCapabilityUsecases
}

func (f fakeDisableCapabilityUsecase) Execute(ctx context.Context, key string, reason string) error {
	return f.fake.ExecuteDisable(ctx, key, reason)
}

func newCapabilityHandlerForTest(t *testing.T) (*CapabilityHandler, *fakeCapabilityUsecases) {
	t.Helper()

	capability, err := domain.NewCapability(domain.Capability{
		Key:                 "capability.manage",
		Domain:              "capability",
		Operation:           "manage",
		Method:              "*",
		Path:                "/api/admin/capabilities",
		DefaultEnabled:      true,
		Enabled:             true,
		RequiredPermission:  "capability.manage",
		RiskLevel:           domain.RiskHigh,
		AuditRequired:       true,
		IdempotencyRequired: false,
		OwnerPackage:        "internal/modules/capability/transport/http",
		TestProof:           "handler tests and SQL proof placeholders",
	})
	if err != nil {
		t.Fatalf("NewCapability() error = %v", err)
	}

	fake := &fakeCapabilityUsecases{
		capabilities: map[string]domain.Capability{
			capability.Key: capability,
		},
	}

	handler := NewCapabilityHandler(
		fake,
		fakeShowCapabilityUsecase{fake: fake},
		fakeEnableCapabilityUsecase{fake: fake},
		fakeDisableCapabilityUsecase{fake: fake},
	)

	return handler, fake
}

type testEnvelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type capabilityResponseForTest struct {
	Key            string `json:"key"`
	RiskLevel      string `json:"risk_level"`
	Enabled        bool   `json:"enabled"`
	DisabledReason string `json:"disabled_reason"`
}

func decodeResponse(t *testing.T, body string, out any) {
	t.Helper()

	if err := json.Unmarshal([]byte(body), out); err != nil {
		t.Fatalf("json.Unmarshal(%q) error = %v", body, err)
	}
}

func decodeRawData(t *testing.T, raw json.RawMessage, out any) {
	t.Helper()

	if len(raw) == 0 {
		t.Fatal("data is empty")
	}
	if err := json.Unmarshal(raw, out); err != nil {
		t.Fatalf("json.Unmarshal(data) error = %v", err)
	}
}

var _ ListCapabilitiesUsecase = (*fakeCapabilityUsecases)(nil)
var _ ShowCapabilityUsecase = fakeShowCapabilityUsecase{}
var _ EnableCapabilityUsecase = fakeEnableCapabilityUsecase{}
var _ DisableCapabilityUsecase = fakeDisableCapabilityUsecase{}

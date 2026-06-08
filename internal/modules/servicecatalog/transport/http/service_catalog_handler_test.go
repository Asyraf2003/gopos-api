package http

import (
	"context"
	"encoding/json"
	stdhttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	servicecatalogusecase "pos-go/internal/modules/servicecatalog/usecase"

	"github.com/labstack/echo/v4"
)

type listServiceCatalogItemsFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
) ([]servicecatalogusecase.ServiceCatalogItemResult, error)

func (f listServiceCatalogItemsFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
) ([]servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type lookupServiceCatalogItemsFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
) ([]servicecatalogusecase.ServiceCatalogLookupResult, error)

func (f lookupServiceCatalogItemsFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
) ([]servicecatalogusecase.ServiceCatalogLookupResult, error) {
	return f(ctx, cmd)
}

type showServiceCatalogItemFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error)

func (f showServiceCatalogItemFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type createServiceCatalogItemFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error)

func (f createServiceCatalogItemFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type updateServiceCatalogItemFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error)

func (f updateServiceCatalogItemFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type activateServiceCatalogItemFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.ActivateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error)

func (f activateServiceCatalogItemFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.ActivateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

type deactivateServiceCatalogItemFunc func(
	ctx context.Context,
	cmd servicecatalogusecase.DeactivateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error)

func (f deactivateServiceCatalogItemFunc) Execute(
	ctx context.Context,
	cmd servicecatalogusecase.DeactivateServiceCatalogItemCommand,
) (servicecatalogusecase.ServiceCatalogItemResult, error) {
	return f(ctx, cmd)
}

func TestServiceCatalogHandler_ListParsesQueryAndReturnsEnvelope(t *testing.T) {
	h := newTestServiceCatalogHandler(t)

	var got servicecatalogusecase.ListServiceCatalogItemsCommand
	h.list = listServiceCatalogItemsFunc(func(
		ctx context.Context,
		cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogItemResult, error) {
		_ = ctx
		got = cmd

		return []servicecatalogusecase.ServiceCatalogItemResult{
			testServiceCatalogItemResult(),
		}, nil
	})

	c, rec := newServiceCatalogTestContext(
		stdhttp.MethodGet,
		"/items?q=wash&status=active&page=2&per_page=5",
		"",
	)

	if err := h.List(c); err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if got.Query != "wash" {
		t.Fatalf("query = %q, want wash", got.Query)
	}
	if string(got.Status) != "active" {
		t.Fatalf("status = %q, want active", got.Status)
	}
	if got.Page != 2 {
		t.Fatalf("page = %d, want 2", got.Page)
	}
	if got.PerPage != 5 {
		t.Fatalf("per_page = %d, want 5", got.PerPage)
	}

	assertJSONStatus(t, rec, stdhttp.StatusOK)
	assertSuccessEnvelope(t, rec)
}

func TestServiceCatalogHandler_LookupParsesActiveOnly(t *testing.T) {
	h := newTestServiceCatalogHandler(t)

	var got servicecatalogusecase.LookupServiceCatalogItemsCommand
	h.lookup = lookupServiceCatalogItemsFunc(func(
		ctx context.Context,
		cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
	) ([]servicecatalogusecase.ServiceCatalogLookupResult, error) {
		_ = ctx
		got = cmd

		return []servicecatalogusecase.ServiceCatalogLookupResult{
			{
				ID:                 "svc_1",
				Name:               "Express Wash",
				DefaultPriceRupiah: 15000,
			},
		}, nil
	})

	c, rec := newServiceCatalogTestContext(
		stdhttp.MethodGet,
		"/items/lookup?q=wash&limit=7&active_only=true",
		"",
	)

	if err := h.Lookup(c); err != nil {
		t.Fatalf("Lookup() error = %v", err)
	}

	if got.Query != "wash" {
		t.Fatalf("query = %q, want wash", got.Query)
	}
	if got.Limit != 7 {
		t.Fatalf("limit = %d, want 7", got.Limit)
	}
	if got.IncludeInactive {
		t.Fatal("include_inactive = true, want false when active_only=true")
	}

	assertJSONStatus(t, rec, stdhttp.StatusOK)
	assertSuccessEnvelope(t, rec)
}

func TestServiceCatalogHandler_CreateRejectsInvalidBodyBeforeUsecase(t *testing.T) {
	h := newTestServiceCatalogHandler(t)

	called := false
	h.create = createServiceCatalogItemFunc(func(
		ctx context.Context,
		cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error) {
		_ = ctx
		_ = cmd
		called = true

		return servicecatalogusecase.ServiceCatalogItemResult{}, nil
	})

	c, _ := newServiceCatalogTestContext(
		stdhttp.MethodPost,
		"/items",
		`{"name":`,
	)

	err := h.Create(c)
	assertHTTPErrorCode(t, err, stdhttp.StatusBadRequest)

	if called {
		t.Fatal("create usecase was called for invalid body")
	}
}

func TestServiceCatalogHandler_UpdateMapsIDAndBody(t *testing.T) {
	h := newTestServiceCatalogHandler(t)

	var got servicecatalogusecase.UpdateServiceCatalogItemCommand
	h.update = updateServiceCatalogItemFunc(func(
		ctx context.Context,
		cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error) {
		_ = ctx
		got = cmd

		result := testServiceCatalogItemResult()
		result.ID = cmd.ID
		result.Name = cmd.Name
		result.DefaultPriceRupiah = cmd.DefaultPriceRupiah

		return result, nil
	})

	c, rec := newServiceCatalogTestContext(
		stdhttp.MethodPut,
		"/items/svc_1",
		`{"name":"Premium Wash","default_price_rupiah":25000}`,
	)
	c.SetParamNames("id")
	c.SetParamValues("svc_1")

	if err := h.Update(c); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if got.ID != "svc_1" {
		t.Fatalf("id = %q, want svc_1", got.ID)
	}
	if got.Name != "Premium Wash" {
		t.Fatalf("name = %q, want Premium Wash", got.Name)
	}
	if got.DefaultPriceRupiah != 25000 {
		t.Fatalf("default_price_rupiah = %d, want 25000", got.DefaultPriceRupiah)
	}

	assertJSONStatus(t, rec, stdhttp.StatusOK)
	assertSuccessEnvelope(t, rec)
}

func TestServiceCatalogHandler_ShowMapsNotFoundTo404(t *testing.T) {
	h := newTestServiceCatalogHandler(t)

	h.show = showServiceCatalogItemFunc(func(
		ctx context.Context,
		cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
	) (servicecatalogusecase.ServiceCatalogItemResult, error) {
		_ = ctx
		if cmd.ID != "svc_missing" {
			t.Fatalf("id = %q, want svc_missing", cmd.ID)
		}

		return servicecatalogusecase.ServiceCatalogItemResult{},
			servicecatalogusecase.ErrServiceCatalogItemNotFound
	})

	c, _ := newServiceCatalogTestContext(
		stdhttp.MethodGet,
		"/items/svc_missing",
		"",
	)
	c.SetParamNames("id")
	c.SetParamValues("svc_missing")

	err := h.Show(c)
	assertHTTPErrorCode(t, err, stdhttp.StatusNotFound)
}

func newTestServiceCatalogHandler(t *testing.T) ServiceCatalogHandler {
	t.Helper()

	failItem := func(name string) func() (servicecatalogusecase.ServiceCatalogItemResult, error) {
		return func() (servicecatalogusecase.ServiceCatalogItemResult, error) {
			t.Fatalf("%s usecase should not be called", name)
			return servicecatalogusecase.ServiceCatalogItemResult{}, nil
		}
	}

	failItems := func(name string) func() ([]servicecatalogusecase.ServiceCatalogItemResult, error) {
		return func() ([]servicecatalogusecase.ServiceCatalogItemResult, error) {
			t.Fatalf("%s usecase should not be called", name)
			return nil, nil
		}
	}

	failLookups := func(name string) func() ([]servicecatalogusecase.ServiceCatalogLookupResult, error) {
		return func() ([]servicecatalogusecase.ServiceCatalogLookupResult, error) {
			t.Fatalf("%s usecase should not be called", name)
			return nil, nil
		}
	}

	return NewServiceCatalogHandler(
		listServiceCatalogItemsFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.ListServiceCatalogItemsCommand,
		) ([]servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItems("list")()
		}),
		lookupServiceCatalogItemsFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.LookupServiceCatalogItemsCommand,
		) ([]servicecatalogusecase.ServiceCatalogLookupResult, error) {
			_ = ctx
			_ = cmd
			return failLookups("lookup")()
		}),
		showServiceCatalogItemFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.ShowServiceCatalogItemCommand,
		) (servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItem("show")()
		}),
		createServiceCatalogItemFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.CreateServiceCatalogItemCommand,
		) (servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItem("create")()
		}),
		updateServiceCatalogItemFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.UpdateServiceCatalogItemCommand,
		) (servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItem("update")()
		}),
		activateServiceCatalogItemFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.ActivateServiceCatalogItemCommand,
		) (servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItem("activate")()
		}),
		deactivateServiceCatalogItemFunc(func(
			ctx context.Context,
			cmd servicecatalogusecase.DeactivateServiceCatalogItemCommand,
		) (servicecatalogusecase.ServiceCatalogItemResult, error) {
			_ = ctx
			_ = cmd
			return failItem("deactivate")()
		}),
	)
}

func newServiceCatalogTestContext(
	method string,
	target string,
	body string,
) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	var reader *strings.Reader
	if body == "" {
		reader = strings.NewReader("")
	} else {
		reader = strings.NewReader(body)
	}

	req := httptest.NewRequest(method, target, reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func testServiceCatalogItemResult() servicecatalogusecase.ServiceCatalogItemResult {
	now := time.Date(2026, 6, 8, 10, 30, 0, 0, time.UTC)

	return servicecatalogusecase.ServiceCatalogItemResult{
		ID:                 "svc_1",
		Name:               "Express Wash",
		NormalizedName:     "express wash",
		DefaultPriceRupiah: 15000,
		IsActive:           true,
		Status:             "active",
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

func assertJSONStatus(
	t *testing.T,
	rec *httptest.ResponseRecorder,
	want int,
) {
	t.Helper()

	if rec.Code != want {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, want, rec.Body.String())
	}

	if got := rec.Header().Get(echo.HeaderContentType); !strings.Contains(got, echo.MIMEApplicationJSON) {
		t.Fatalf("content-type = %q, want JSON", got)
	}
}

func assertSuccessEnvelope(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()

	var body struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data"`
		Meta    json.RawMessage `json:"meta"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v; body = %s", err, rec.Body.String())
	}

	if !body.Success {
		t.Fatalf("success = false, want true; body = %s", rec.Body.String())
	}
	if len(body.Data) == 0 || string(body.Data) == "null" {
		t.Fatalf("data is empty; body = %s", rec.Body.String())
	}
	if len(body.Meta) == 0 || string(body.Meta) == "null" {
		t.Fatalf("meta is empty; body = %s", rec.Body.String())
	}
}

func assertHTTPErrorCode(t *testing.T, err error, want int) {
	t.Helper()

	if err == nil {
		t.Fatalf("error = nil, want HTTP %d", want)
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}

	if httpErr.Code != want {
		t.Fatalf("status = %d, want %d", httpErr.Code, want)
	}
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"pos-go/internal/modules/auth/domain"

	"github.com/labstack/echo/v4"
)

func newAuthzTestContext(principal *domain.Principal) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if principal != nil {
		req = req.WithContext(WithPrincipal(req.Context(), *principal))
	}

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func assertHTTPErrorCode(t *testing.T, err error, want int) {
	t.Helper()

	if err == nil {
		t.Fatal("handler() error = nil, want error")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("error type = %T, want *echo.HTTPError", err)
	}
	if httpErr.Code != want {
		t.Fatalf("status = %d, want %d", httpErr.Code, want)
	}
}

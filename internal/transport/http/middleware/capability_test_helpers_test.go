package middleware

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func newCapabilityTestContext() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/test", nil)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

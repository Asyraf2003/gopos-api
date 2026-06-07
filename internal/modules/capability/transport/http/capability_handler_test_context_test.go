package http

import (
	"io"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

func newCapabilityContext(
	e *echo.Echo,
	method string,
	target string,
	body string,
) (echo.Context, *httptest.ResponseRecorder) {
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}

	req := httptest.NewRequest(method, target, reader)
	if body != "" {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

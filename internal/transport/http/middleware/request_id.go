package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/labstack/echo/v4"
)

const requestIDKey = "request_id"

func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := c.Request().Header.Get(echo.HeaderXRequestID)
		if requestID == "" {
			requestID = newRequestID()
		}

		c.Set(requestIDKey, requestID)
		c.Response().Header().Set(echo.HeaderXRequestID, requestID)

		return next(c)
	}
}

func newRequestID() string {
	buf := make([]byte, 16)
	_, _ = rand.Read(buf)

	return hex.EncodeToString(buf)
}

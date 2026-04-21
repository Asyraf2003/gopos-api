package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "internal server error",
				})
			}
		}()

		return next(c)
	}
}

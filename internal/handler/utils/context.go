package utils

import (
	"context"

	"github.com/labstack/echo"
)

// EchoToContext transform echo.Context to context.Context
func EchoToContext(c echo.Context, key string) context.Context {
	var keyInterface interface{}
	keyInterface = key
	return context.WithValue(c.Request().Context(), keyInterface, c.Get(key))
}

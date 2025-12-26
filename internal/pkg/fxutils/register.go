package fxutils

import (
	"github.com/labstack/echo/v4"
)

type Registrar interface {
	Register(e *echo.Echo)
}

func Register[T Registrar]() func(e *echo.Echo, c T) {
	return func(e *echo.Echo, c T) {
		c.Register(e)
	}
}

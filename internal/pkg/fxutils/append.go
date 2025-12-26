package fxutils

import (
	"context"
	"go.uber.org/fx"
)

type Component interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func Append[T Component]() func(lc fx.Lifecycle, c T) {
	return func(lc fx.Lifecycle, c T) {
		lc.Append(fx.Hook{
			OnStart: c.Start,
			OnStop:  c.Stop,
		})
	}
}

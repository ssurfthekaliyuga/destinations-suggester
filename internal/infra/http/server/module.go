package server

import (
	"destinations-suggester/internal/infra/http/server/handlers/suggestions"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module("http.server",
	fx.Provide(
		echo.New, New,
	),
	fx.Provide(
		suggestions.NewLister,
	),
	fx.Invoke(
		func(h *suggestions.Lister, e *echo.Echo) {
			h.Register(e)
		},
	),
	fx.Invoke(
		func(lc fx.Lifecycle, s *Server) {
			lc.Append(fx.Hook{
				OnStart: s.Start,
				OnStop:  s.Stop,
			})
		},
	),
)

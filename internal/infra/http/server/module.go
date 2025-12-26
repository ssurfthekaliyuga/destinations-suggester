package server

import (
	"context"
	"destinations-suggester/internal/infra/http/server/handlers/suggestions"
	"destinations-suggester/internal/pkg/fxutils"
	"destinations-suggester/internal/pkg/sl"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module("http.server",
	fx.Provide(
		echo.New, New,
	),
	fx.Provide(
		suggestions.NewLister,
	),
	fx.Invoke(
		fxutils.Register[*suggestions.Lister](),
	),
	fx.Invoke(
		func(lc fx.Lifecycle, sh fx.Shutdowner, srv *Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := srv.Start(ctx); err != nil {
							slog.Error("cannot start http server", sl.Error(err))
							if err := sh.Shutdown(); err != nil {
								slog.Error("cannot shutdown application", sl.Error(err))
							}
						}
					}()

					return nil
				},
				OnStop: srv.Stop,
			})
		},
	),
)

package server

import (
	"destinations-suggester/internal/infra/http/server/handlers/suggestions"
	"destinations-suggester/internal/pkg/fxutils"
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
		fxutils.Register[*suggestions.Lister](),
	),
	fx.Invoke(
		fxutils.Append[*Server](),
	),
)

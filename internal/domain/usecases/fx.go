package usecases

import (
	"destinations-suggester/internal/domain/usecases/places"
	"destinations-suggester/internal/domain/usecases/suggestions"
	"go.uber.org/fx"
)

var Fx = fx.Module("usecases", fx.Provide(
	places.NewRideEventsHandler,
	places.NewSearchEventsHandler,
	suggestions.NewCalculator,
	suggestions.NewLister,
))

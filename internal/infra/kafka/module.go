package kafka

import (
	"destinations-suggester/internal/infra/kafka/consumers"
	"destinations-suggester/internal/pkg/fxutils"
	"go.uber.org/fx"
)

var Module = fx.Module("kafka",
	fx.Provide(
		consumers.NewRideEvents,
		consumers.NewSearchEvents,
	),
	fx.Invoke(
		fxutils.Append[*consumers.RideEvents](),
		fxutils.Append[*consumers.SearchEvents](),
	),
)

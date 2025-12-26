package kafka

import (
	"destinations-suggester/internal/infra/kafka/consumers"
	"go.uber.org/fx"
)

var Module = fx.Module("kafka",
	fx.Provide(
		consumers.NewRideEvents,
		consumers.NewSearchEvents,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, events *consumers.RideEvents) {
			lc.Append(fx.Hook{
				OnStart: events.Start,
				OnStop:  events.Stop,
			})
		},
		func(lc fx.Lifecycle, events *consumers.SearchEvents) {
			lc.Append(fx.Hook{
				OnStart: events.Start,
				OnStop:  events.Stop,
			})
		},
	),
)

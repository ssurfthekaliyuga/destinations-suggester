package consumers

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"destinations-suggester/internal/pkg/sl"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type rideEventsHandler interface {
	Handle(ctx context.Context, event *places.Ride) error
}

type RideEvents struct {
	logger *slog.Logger
	reader *kafka.Reader
	events rideEventsHandler
}

func NewRideEvents(
	reader *kafka.Reader,
	eventsHandler rideEventsHandler,
) *RideEvents {
	return &RideEvents{
		logger: sl.WithComponent("kafka.consumers.RideEvents"),
		reader: reader,
		events: eventsHandler,
	}
}

func (c *RideEvents) Start(ctx context.Context) error { // todo add validation and dead letter queue
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			c.logger.Error("cannot fetch message", sl.Error(err))
			continue
		}

		logger := c.logger.With(
			slog.String("key", string(msg.Key)),
			slog.Int("partition", msg.Partition),
			slog.Int64("offset", msg.Offset),
		)

		var event places.Ride
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			logger.Error("cannot unmarshal ride event", sl.Error(err))
			continue
		}

		if err := c.events.Handle(ctx, &event); err != nil {
			logger.Error("cannot handle ride event", sl.Error(err))
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			logger.Error("cannot commit message", sl.Error(err))
			continue
		}
	}
}

func (c *RideEvents) Stop(_ context.Context) error {
	return c.reader.Close()
}

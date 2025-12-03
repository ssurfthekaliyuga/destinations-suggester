package consumers

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type rideEventsHandler interface {
	Handle(ctx context.Context, event *places.Ride) error
}

type RideEvents struct {
	reader *kafka.Reader
	events rideEventsHandler
	errors errorsHandler
}

func NewRideEvents(
	reader *kafka.Reader,
	eventsHandler rideEventsHandler,
	errorsHandler errorsHandler,
) *RideEvents {
	return &RideEvents{
		reader: reader,
		events: eventsHandler,
		errors: errorsHandler,
	}
}

func (c *RideEvents) StartConsuming(ctx context.Context) error { // todo add validation and dead letter queue
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			c.errors.Handle(ctx, "cannot fetch message", err)
			continue
		}

		var event places.Ride
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.errors.Handle(ctx, "cannot unmarshal ride event", err)
			continue
		}

		if err := c.events.Handle(ctx, &event); err != nil {
			c.errors.Handle(ctx, "cannot handle ride event", err)
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			c.errors.Handle(ctx, "cannot commit message", err)
			continue
		}
	}
}

func (c *RideEvents) Close() error {
	return c.reader.Close()
}

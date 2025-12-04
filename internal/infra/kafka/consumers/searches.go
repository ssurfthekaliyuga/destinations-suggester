package consumers

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type searchEventsHandler interface {
	Handle(ctx context.Context, event *places.Search) error
}

type SearchEvents struct {
	reader *kafka.Reader
	events searchEventsHandler
	errors errorsHandler
}

func NewSearchEvents(
	reader *kafka.Reader,
	eventsHandler searchEventsHandler,
	errorsHandler errorsHandler,
) *SearchEvents {
	return &SearchEvents{
		reader: reader,
		events: eventsHandler,
		errors: errorsHandler,
	}
}

func (c *SearchEvents) StartConsuming(ctx context.Context) error { // todo add validation and dead letter queue
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

		var event places.Search
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.errors.Handle(ctx, "cannot unmarshal search event", err)
			continue
		}

		if err := c.events.Handle(ctx, &event); err != nil {
			c.errors.Handle(ctx, "cannot handle search event", err)
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			c.errors.Handle(ctx, "cannot commit message", err)
			continue
		}
	}
}

func (c *SearchEvents) Close() error {
	return c.reader.Close()
}

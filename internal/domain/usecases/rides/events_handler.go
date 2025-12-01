package rides

import (
	"context"
	"destinations-suggester/internal/domain/models/rides"
	"fmt"
	"github.com/google/uuid"
)

type rideEvents interface {
	Save(ctx context.Context, event *rides.Event) error
}

type rideSuggestions interface {
	DoTasks(ctx context.Context, userID uuid.UUID) error
}

type EventsHandler struct {
	rideEvents      rideEvents
	rideSuggestions rideSuggestions
}

func NewEventsHandler(
	rideEvents rideEvents,
	rideSuggestions rideSuggestions,
) *EventsHandler {
	return &EventsHandler{
		rideEvents:      rideEvents,
		rideSuggestions: rideSuggestions,
	}
}

func (h *EventsHandler) Handle(ctx context.Context, event *rides.Event) error {
	if err := h.rideEvents.Save(ctx, event); err != nil {
		return fmt.Errorf("cannot save ride event: %w", err)
	}
	if err := h.rideSuggestions.DoTasks(ctx, event.UserID); err != nil {
		return fmt.Errorf("cannot create calculate suggestion task: %w", err)
	}
	return nil
}

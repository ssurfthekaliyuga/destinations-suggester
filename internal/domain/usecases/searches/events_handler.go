package searches

import (
	"context"
	"destinations-suggester/internal/domain/models/searches"
	"fmt"
	"github.com/google/uuid"
)

type searchEvents interface {
	Save(ctx context.Context, event *searches.Event) error
}

type rideSuggestions interface {
	DoTasks(ctx context.Context, userID uuid.UUID) error
}

type EventsHandler struct {
	searchEvents    searchEvents
	rideSuggestions rideSuggestions
}

func NewEventsHandler(
	searchEvents searchEvents,
	rideSuggestions rideSuggestions,
) *EventsHandler {
	return &EventsHandler{
		searchEvents:    searchEvents,
		rideSuggestions: rideSuggestions,
	}
}

func (h *EventsHandler) Handle(ctx context.Context, event *searches.Event) error {
	if err := h.searchEvents.Save(ctx, event); err != nil {
		return fmt.Errorf("cannot save ride event: %w", err)
	}
	if err := h.rideSuggestions.DoTasks(ctx, event.UserID); err != nil {
		return fmt.Errorf("cannot create calculate suggestion task: %w", err)
	}
	return nil
}

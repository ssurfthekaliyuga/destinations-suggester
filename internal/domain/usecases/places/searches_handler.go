package places

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"fmt"
)

type searchesSaver interface {
	SaveSearch(ctx context.Context, search *places.Search) error
}

type SearchEventsHandler struct {
	searchesSaver         searchesSaver
	suggestionsCalculator suggestionsCalculator
}

func NewSearchEventsHandler(
	searchesSaver searchesSaver,
	suggestionsCalculator suggestionsCalculator,
) *SearchEventsHandler {
	return &SearchEventsHandler{
		searchesSaver:         searchesSaver,
		suggestionsCalculator: suggestionsCalculator,
	}
}

func (h *SearchEventsHandler) Handle(ctx context.Context, search *places.Search) error {
	if err := h.searchesSaver.SaveSearch(ctx, search); err != nil {
		return fmt.Errorf("cannot save search event: %w", err)
	}
	if err := h.suggestionsCalculator.Calculate(ctx, search.UserID); err != nil {
		return fmt.Errorf("cannot create calculate suggestion task: %w", err)
	}
	return nil
}

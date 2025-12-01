package places

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"fmt"
)

type ridesSaver interface {
	Save(ctx context.Context, ride *places.Ride) error
}

type RideEventsHandler struct {
	ridesSaver            ridesSaver
	suggestionsCalculator suggestionsCalculator
}

func NewRideEventsHandler(
	ridesSaver ridesSaver,
	suggestionsCalculator suggestionsCalculator,
) *RideEventsHandler {
	return &RideEventsHandler{
		ridesSaver:            ridesSaver,
		suggestionsCalculator: suggestionsCalculator,
	}
}

func (h *RideEventsHandler) Handle(ctx context.Context, ride *places.Ride) error {
	if err := h.ridesSaver.Save(ctx, ride); err != nil {
		return fmt.Errorf("cannot save ride event: %w", err)
	}
	if err := h.suggestionsCalculator.Calculate(ctx, ride.UserID); err != nil {
		return fmt.Errorf("cannot create calculate suggestion task: %w", err)
	}
	return nil
}

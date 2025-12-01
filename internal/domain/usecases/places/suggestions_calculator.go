package places

import (
	"context"
	"github.com/google/uuid"
)

type suggestionsCalculator interface {
	Calculate(ctx context.Context, userID uuid.UUID) error
}

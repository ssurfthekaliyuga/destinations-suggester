package suggestions

import (
	"destinations-suggester/internal/domain/models/places"
	"github.com/google/uuid"
)

type Query struct {
	UserID                       uuid.UUID
	UserCurrentLocation          places.Coordinates
	ExcludeCurrentLocationRadius int
	Limit                        int
}

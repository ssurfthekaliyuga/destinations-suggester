package rides

import (
	"destinations-suggester/internal/domain/models/places"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID     uuid.UUID
	UserID uuid.UUID
	From   places.Place
	To     places.Place
	Time   time.Time
}

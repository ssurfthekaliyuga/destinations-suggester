package searches

import (
	"destinations-suggester/internal/domain/models/places"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Place  places.Place
	Time   time.Time
}

package places

import (
	"github.com/google/uuid"
	"time"
)

type Search struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Place  Place
	Time   time.Time
}

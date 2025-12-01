package places

import (
	"github.com/google/uuid"
	"time"
)

type Ride struct {
	ID     uuid.UUID
	UserID uuid.UUID
	From   Place
	To     Place
	Time   time.Time
}

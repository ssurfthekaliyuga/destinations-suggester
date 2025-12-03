package suggestions

import (
	"destinations-suggester/internal/domain/models/places"
	"time"
)

type CalculatorConfig struct {
	Params          places.CalculateScoreParams
	NoTasksDelay    time.Duration
	ReleaseTimeout  time.Duration
	MaxWorkers      int
	UserPlacesLimit int
}

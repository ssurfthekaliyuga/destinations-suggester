package suggestions

import (
	"destinations-suggester/internal/domain/models/suggestions"
	"time"
)

type CalculatorConfig struct {
	Params          suggestions.CalculateParams
	NoTasksDelay    time.Duration
	ReleaseTimeout  time.Duration
	MaxWorkers      int
	UserPlacesLimit int
}

package suggestions

import (
	"time"
)

type CalculateParams struct {
	TimeDecayRate     float64
	Now               time.Time
	FreshSearchWindow time.Duration
	FreshSearchWeight float64
	StaleSearchWeight float64
}

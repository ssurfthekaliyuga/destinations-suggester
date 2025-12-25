package suggestions

import (
	"destinations-suggester/internal/domain/models/places"
)

type CalculatorConfig struct {
	Params          places.CalculateScoreParams
	UserPlacesLimit int
}

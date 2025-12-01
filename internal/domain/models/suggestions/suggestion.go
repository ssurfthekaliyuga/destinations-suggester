package suggestions

import (
	"destinations-suggester/internal/domain/models/places"
)

type Suggestion struct {
	Place places.Place
	Score float64
}

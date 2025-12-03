package suggestions

import (
	"destinations-suggester/internal/domain/models/places"
	"github.com/google/uuid"
	"math"
	"time"
)

type UserPlacesStatsQuery struct {
	UserID uuid.UUID
	Limit  int
}

type UserPlaceStat struct {
	Place     places.Place
	UserID    uuid.UUID
	RidesFrom []time.Time
	RidesTo   []time.Time
	Searches  []time.Time
}

func (s *UserPlaceStat) CalculateSuggestion(p *CalculateParams) *Suggestion {
	var score float64
	for _, rideFromTime := range s.RidesFrom {
		score += math.Exp(-p.TimeDecayRate * p.Now.Sub(rideFromTime).Seconds())
	}
	for _, rideToTime := range s.RidesTo {
		score += math.Exp(-p.TimeDecayRate * p.Now.Sub(rideToTime).Seconds())
	}

	for _, searchTime := range s.Searches {
		if p.Now.Sub(searchTime) <= p.FreshSearchWindow {
			score += p.FreshSearchWeight
		} else {
			score += p.StaleSearchWeight
		}
	}

	return &Suggestion{
		Place: s.Place,
		Score: score,
	}
}

package places

import (
	"github.com/google/uuid"
	"math"
	"time"
)

type UserStatsQuery struct {
	UserID uuid.UUID
	Limit  int
}

type UserStat struct {
	Place     Place
	UserID    uuid.UUID
	RidesFrom []time.Time
	RidesTo   []time.Time
	Searches  []time.Time
}

func (s *UserStat) Score(p *CalculateScoreParams) float64 {
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

	return score
}

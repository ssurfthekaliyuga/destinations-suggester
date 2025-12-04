package shared

import "destinations-suggester/internal/domain/models/suggestions"

type Suggestion struct {
	Place Place   `json:"place"`
	Score float64 `json:"score"`
}

func (Suggestion) FromModel(suggestion *suggestions.Suggestion) Suggestion {
	return Suggestion{
		Place: Place{}.FromModel(suggestion.Place),
		Score: suggestion.Score,
	}
}

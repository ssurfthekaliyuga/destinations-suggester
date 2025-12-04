package suggestions

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"destinations-suggester/internal/domain/models/suggestions"
	"fmt"
	"github.com/google/uuid"
)

type suggestionsLister interface {
	List(ctx context.Context, query *suggestions.Query) ([]suggestions.Suggestion, error)
}

type Lister struct {
	conf             *ListerConfig
	suggestionsQuery suggestionsLister
}

func NewLister(
	conf *ListerConfig,
	suggestionsQuery suggestionsLister,
) (*Lister, error) {
	return &Lister{
		conf:             conf,
		suggestionsQuery: suggestionsQuery,
	}, nil
}

func (p *Lister) List(ctx context.Context, userID uuid.UUID, userLocations places.Coordinates) ([]suggestions.Suggestion, error) {
	suggestionsSlice, err := p.suggestionsQuery.List(ctx, &suggestions.Query{
		UserID:                       userID,
		UserCurrentLocation:          userLocations,
		ExcludeCurrentLocationRadius: p.conf.ExcludeCurrentLocationRadius,
		Limit:                        p.conf.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot list user places suggestionsSlice: %w", err)
	}

	return suggestionsSlice, nil
}

package suggestions

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"destinations-suggester/internal/domain/models/suggestions"
	"fmt"
	"github.com/google/uuid"
)

type suggestionsQuery interface {
	List(ctx context.Context, query *suggestions.Query) ([]suggestions.Suggestion, error)
}

type Provider struct {
	conf             *ProviderConfig
	suggestionsQuery suggestionsQuery
}

func NewProvider(
	conf *ProviderConfig,
	suggestionsQuery suggestionsQuery,
) (*Provider, error) {
	return &Provider{
		conf:             conf,
		suggestionsQuery: suggestionsQuery,
	}, nil
}

func (p *Provider) List(ctx context.Context, userID uuid.UUID, userLocations places.Coordinates) ([]suggestions.Suggestion, error) {
	res, err := p.suggestionsQuery.List(ctx, &suggestions.Query{
		UserID:                       userID,
		UserCurrentLocation:          userLocations,
		ExcludeCurrentLocationRadius: p.conf.ExcludeCurrentLocationRadius,
		Limit:                        p.conf.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot list user places suggestions: %w", err)
	}

	return res, nil
}

package suggestions

import (
	"context"
	"destinations-suggester/internal/domain/models/suggestions"
	"fmt"
	"github.com/google/uuid"
)

type errorsHandler interface {
	Handle(ctx context.Context, msg string, err error)
}

type userPlacesQuery interface {
	ListStats(ctx context.Context, query *suggestions.UserStatsQuery) ([]suggestions.UserStat, error)
}

type suggestionsRepo interface {
	ClaimLastTask(ctx context.Context, userID uuid.UUID) error
	Save(ctx context.Context, userID uuid.UUID, suggestions []suggestions.UserStat) error
}

type Calculator struct {
	conf            *CalculatorConfig
	userPlacesQuery userPlacesQuery
	suggestionsRepo suggestionsRepo
	errorsHandler   errorsHandler
}

func NewCalculator(
	conf *CalculatorConfig,
	userPlacesQuery userPlacesQuery,
	suggestionsRepo suggestionsRepo,
	errorsHandler errorsHandler,
) *Calculator {
	return &Calculator{
		conf:            conf,
		userPlacesQuery: userPlacesQuery,
		suggestionsRepo: suggestionsRepo,
		errorsHandler:   errorsHandler,
	}
}

func (c *Calculator) DoTasks(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

	}
}

func (c *Calculator) doTask(ctx context.Context, task *suggestions.CalculateTask) error {
	stats, err := c.userPlacesQuery.ListStats(ctx, &suggestions.UserStatsQuery{
		UserID: task.UserID,
		Limit:  c.conf.UserPlacesLimit,
	})
	if err != nil {
		return fmt.Errorf("cannot list user places stats: %w", err)
	}

	suggestionsSlice := make([]suggestions.Suggestion, 0, len(stats))
	for _, stat := range stats {
		suggestionsSlice = append(suggestionsSlice, suggestions.Suggestion{
			Place: stat.Place,
			Score: 0,
		})
	}

	c.suggestionsRepo.ClaimLastTask(ctx, task.UserID)
}

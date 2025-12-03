package suggestions

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"destinations-suggester/internal/domain/models/suggestions"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type errorsHandler interface { // todo there is double
	Handle(ctx context.Context, msg string, err error)
}

type userPlacesLister interface {
	ListUserStats(ctx context.Context, query *places.UserStatsQuery) ([]places.UserStat, error)
}

type suggestionsRepo interface {
	Save(ctx context.Context, userID uuid.UUID, suggestions []suggestions.Suggestion) error
	CreateCalculateTask(ctx context.Context, userID uuid.UUID) error
	ClaimUserLastCalculateTask(ctx context.Context) (*suggestions.CalculateTask, error)
	UpdateCalculateTasksByUserID(ctx context.Context, userID uuid.UUID, fn suggestions.UpdateCalculateTaskFn) error
}

type Calculator struct {
	conf             *CalculatorConfig
	userPlacesLister userPlacesLister
	suggestionsRepo  suggestionsRepo
	errorsHandler    errorsHandler
}

func NewCalculator(
	conf *CalculatorConfig,
	userPlacesLister userPlacesLister,
	suggestionsRepo suggestionsRepo,
	errorsHandler errorsHandler,
) *Calculator {
	return &Calculator{
		conf:             conf,
		userPlacesLister: userPlacesLister,
		suggestionsRepo:  suggestionsRepo,
		errorsHandler:    errorsHandler,
	}
}

func (c *Calculator) Calculate(ctx context.Context, userID uuid.UUID) error {
	return c.suggestionsRepo.CreateCalculateTask(ctx, userID)
}

func (c *Calculator) StartDoingTasks(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		task, err := c.suggestionsRepo.ClaimUserLastCalculateTask(ctx)
		if errors.Is(err, suggestions.ErrNoTasks) {
			continue
		}
		if err != nil {
			c.errorsHandler.Handle(ctx, "cannot claim user last calculate task", err)
			continue
		}

		if err := c.doTask(ctx, task); err != nil {
			c.errorsHandler.Handle(ctx, "cannot do task", err)
			continue
		}
	}
}

func (c *Calculator) doTask(ctx context.Context, task *suggestions.CalculateTask) error {
	stats, err := c.userPlacesLister.ListUserStats(ctx, &places.UserStatsQuery{
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
			Score: stat.Score(&c.conf.Params),
		})
	}

	if err := c.suggestionsRepo.Save(ctx, task.UserID, suggestionsSlice); err != nil {
		return fmt.Errorf("cannot save suggestions: %w", err)
	}

	updateFn := func(ctx context.Context, existing *suggestions.CalculateTask) *suggestions.CalculateTask {
		if existing.CreatedAt.Before(task.CreatedAt) {
			return existing
		}

		return &suggestions.CalculateTask{
			ID:        existing.ID,
			UserID:    existing.UserID,
			Status:    suggestions.CalculateTaskStatusCompleted,
			CreatedAt: existing.CreatedAt,
		}
	}

	if err := c.suggestionsRepo.UpdateCalculateTasksByUserID(ctx, task.UserID, updateFn); err != nil {
		return fmt.Errorf("cannot update calculate task by user id: %w", err)
	}

	return nil
}

package repositories

import (
	"context"
	"destinations-suggester/internal/domain/models/suggestions"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Suggestions struct {
	pool *pgxpool.Pool
}

func NewSuggestions(pool *pgxpool.Pool) *Suggestions {
	return &Suggestions{
		pool: pool,
	}
}

func (r *Suggestions) Save(ctx context.Context, userID uuid.UUID, suggestions []suggestions.Suggestion) error {
	//TODO implement me
	panic("implement me")
}

func (r *Suggestions) List(ctx context.Context, query *suggestions.Query) ([]suggestions.Suggestion, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Suggestions) CreateCalculateTask(ctx context.Context, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *Suggestions) ClaimUserLastCalculateTask(ctx context.Context) (*suggestions.CalculateTask, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Suggestions) UpdateCalculateTasksByUserID(ctx context.Context, userID uuid.UUID, fn suggestions.UpdateCalculateTaskFn) error {
	//TODO implement me
	panic("implement me")
}

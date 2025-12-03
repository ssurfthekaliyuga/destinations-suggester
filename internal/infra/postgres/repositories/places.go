package repositories

import (
	"context"
	"destinations-suggester/internal/domain/models/places"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Places struct {
	pool *pgxpool.Pool
}

func NewPlaces(pool *pgxpool.Pool) *Places {
	return &Places{
		pool: pool,
	}
}

func (r *Places) ListUserStats(ctx context.Context, query *places.UserStatsQuery) ([]places.UserStat, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Places) SaveSearch(ctx context.Context, search *places.Search) error {
	//TODO implement me
	panic("implement me")
}

func (r *Places) SaveRide(ctx context.Context, ride *places.Ride) error {
	//TODO implement me
	panic("implement me")
}

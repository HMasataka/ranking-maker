package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RankRepository interface {
	Add(ctx context.Context, key string, score float64, member any) error
	Range(ctx context.Context, key string, min, max int64) ([]redis.Z, error)
}

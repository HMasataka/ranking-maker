package repository

import (
	"context"

	"github.com/HMasataka/ranking-maker/domain/entity"
	"github.com/redis/go-redis/v9"
)

type RankRepository interface {
	Add(ctx context.Context, key string, score float64, member any) error
	Range(ctx context.Context, key string, min, max int64) ([]redis.Z, error)
	Rank(ctx context.Context, key string, item *entity.Item) (int64, error)
	RevRank(ctx context.Context, key string, item *entity.Item) (int64, error)
	Delete(ctx context.Context, key string) (int64, error)
}

package service

import (
	"context"

	"github.com/HMasataka/ranking-maker/domain/entity"
	"github.com/HMasataka/ranking-maker/domain/repository"
	"github.com/redis/go-redis/v9"
)

type RankService interface {
	Add(ctx context.Context, key string, score float64, member any) error
	Range(ctx context.Context, key string, min, max int64) ([]redis.Z, error)
	Rank(ctx context.Context, key string, item *entity.Item) (int64, error)
	RevRank(ctx context.Context, key string, item *entity.Item) (int64, error)
}

type rankService struct {
	rankRepository repository.RankRepository
}

func NewRankService(
	rankRepository repository.RankRepository,
) RankService {
	return &rankService{
		rankRepository: rankRepository,
	}
}

func (c *rankService) Add(ctx context.Context, key string, score float64, member any) error {
	return c.rankRepository.Add(ctx, key, score, member)
}

func (c *rankService) Range(ctx context.Context, key string, min, max int64) ([]redis.Z, error) {
	return c.rankRepository.Range(ctx, key, min, max)
}

func (c *rankService) Rank(ctx context.Context, key string, item *entity.Item) (int64, error) {
	return c.rankRepository.Rank(ctx, key, item)
}

func (c *rankService) RevRank(ctx context.Context, key string, item *entity.Item) (int64, error) {
	return c.rankRepository.RevRank(ctx, key, item)
}

package service

import (
	"context"
	"time"

	"github.com/HMasataka/ranking-maker/domain/repository"
	"github.com/HMasataka/transactor"
	"github.com/HMasataka/transactor/redis"
)

type ScoreService interface {
	Add(ctx context.Context, key, member string) error
	Count(ctx context.Context, key string, expired time.Duration) (int64, error)
}

type scoreService struct {
	redisTransactor transactor.Transactor
	scoreRepository repository.ScoreRepository
}

func NewScoreService(
	redisConnectionProvider redis.ConnectionProvider,
	scoreRepository repository.ScoreRepository,
) ScoreService {
	return &scoreService{
		redisTransactor: redis.NewTransactor(redisConnectionProvider),
		scoreRepository: scoreRepository,
	}
}

func (c *scoreService) Count(ctx context.Context, key string, expired time.Duration) (int64, error) {
	var count int64

	if err := c.redisTransactor.Required(ctx, func(ctx context.Context) error {
		var err error
		count, err = c.scoreRepository.Count(ctx, key, expired)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return count, err
	}

	return count, nil
}

func (c *scoreService) Add(ctx context.Context, key, member string) error {
	return c.redisTransactor.Required(ctx, func(ctx context.Context) error {
		return c.scoreRepository.Add(ctx, key, member)
	})
}

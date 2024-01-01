package usecase

import (
	"context"
	"time"

	"github.com/HMasataka/ranking-maker/domain/repository"
	"github.com/HMasataka/transactor"
	"github.com/HMasataka/transactor/redis"
)

type ScoreUseCase interface {
	Increment(ctx context.Context, key, member string) error
	Count(ctx context.Context, key string) (int64, error)
}

type scoreUseCase struct {
	redisTransactor transactor.Transactor
	scoreRepository repository.ScoreRepository
}

func NewScoreUseCase(
	redisConnectionProvider redis.ConnectionProvider,
	scoreRepository repository.ScoreRepository,
) ScoreUseCase {
	return &scoreUseCase{
		redisTransactor: redis.NewTransactor(redisConnectionProvider),
		scoreRepository: scoreRepository,
	}
}

func (c *scoreUseCase) Count(ctx context.Context, key string) (int64, error) {
	var res int64

	if err := c.redisTransactor.Required(ctx, func(ctx context.Context) error {
		count, err := c.scoreRepository.Count(ctx, key, time.Hour)
		if err != nil {
			return err
		}

		res = count

		return nil
	}); err != nil {
		return res, err
	}

	return res, nil
}

func (c *scoreUseCase) Increment(ctx context.Context, key, member string) error {
	return c.redisTransactor.Required(ctx, func(ctx context.Context) error {
		return c.scoreRepository.Increment(ctx, key, member)
	})
}

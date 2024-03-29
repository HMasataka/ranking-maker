package service

import (
	"context"
	"time"

	"github.com/HMasataka/ranking-maker/domain/repository"
)

type ScoreService interface {
	Add(ctx context.Context, key string, member any) error
	Count(ctx context.Context, key string, expired time.Duration) (int64, error)
}

type scoreService struct {
	scoreRepository repository.ScoreRepository
}

func NewScoreService(
	scoreRepository repository.ScoreRepository,
) ScoreService {
	return &scoreService{
		scoreRepository: scoreRepository,
	}
}

func (c *scoreService) Count(ctx context.Context, key string, expired time.Duration) (int64, error) {
	return c.scoreRepository.Count(ctx, key, expired)
}

func (c *scoreService) Add(ctx context.Context, key string, member any) error {
	return c.scoreRepository.Add(ctx, key, member)
}

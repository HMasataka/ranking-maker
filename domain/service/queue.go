package service

import (
	"context"

	"github.com/HMasataka/ranking-maker/domain/repository"
	tx "github.com/HMasataka/transactor/redis"
)

type QueueService interface {
	Push(ctx context.Context, key string, members ...any) error
	Pop(ctx context.Context, key string, count int) ([]string, error)
	Len(ctx context.Context, key string) (int64, error)
}

type queueService struct {
	queueRepository repository.QueueRepository
}

func NewQueueService(
	redisConnectionProvider tx.ConnectionProvider,
	queueRepository repository.QueueRepository,
) QueueService {
	return &queueService{
		queueRepository: queueRepository,
	}
}

func (c *queueService) Push(ctx context.Context, key string, members ...any) error {
	return c.queueRepository.Enqueue(ctx, key, members...)
}

func (c *queueService) Pop(ctx context.Context, key string, count int) ([]string, error) {
	return c.queueRepository.Dequeue(ctx, key, count)
}

func (c *queueService) Len(ctx context.Context, key string) (int64, error) {
	return c.queueRepository.Len(ctx, key)
}

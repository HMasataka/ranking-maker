package service

import (
	"context"

	"github.com/HMasataka/ranking-maker/domain/repository"
)

type QueueService interface {
	Push(ctx context.Context, key string, members ...any) error
	Get(ctx context.Context, key string, start, stop int64) ([]string, error)
	Delete(ctx context.Context, key string, count int64) error
	Pop(ctx context.Context, key string, count int64) ([]string, error)
	Len(ctx context.Context, key string) (int64, error)
}

type queueService struct {
	queueRepository repository.QueueRepository
}

func NewQueueService(
	queueRepository repository.QueueRepository,
) QueueService {
	return &queueService{
		queueRepository: queueRepository,
	}
}

func (c *queueService) Push(ctx context.Context, key string, members ...any) error {
	return c.queueRepository.Enqueue(ctx, key, members...)
}

func (c *queueService) Get(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.queueRepository.Get(ctx, key, start, stop)
}

func (c *queueService) Delete(ctx context.Context, key string, count int64) error {
	_, err := c.queueRepository.Dequeue(ctx, key, count)
	return err
}

func (c *queueService) Pop(ctx context.Context, key string, count int64) ([]string, error) {
	return c.queueRepository.Dequeue(ctx, key, count)
}

func (c *queueService) Len(ctx context.Context, key string) (int64, error) {
	return c.queueRepository.Len(ctx, key)
}

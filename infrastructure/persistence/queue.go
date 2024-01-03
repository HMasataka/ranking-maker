package persistence

import (
	"context"
	"fmt"

	"github.com/HMasataka/ranking-maker/domain/repository"
	tx "github.com/HMasataka/transactor/redis"
)

type queueRepository struct {
	clientProvider tx.ClientProvider
}

func NewQueueRepository(clientProvider tx.ClientProvider) repository.QueueRepository {
	return &queueRepository{
		clientProvider: clientProvider,
	}
}

func (s *queueRepository) getKey(key string) string {
	return fmt.Sprintf("queue:%s", key)
}

func (s *queueRepository) Enqueue(ctx context.Context, key string, members ...any) error {
	_, writer := s.clientProvider.CurrentClient(ctx)

	return writer.LPush(ctx, s.getKey(key), members...).Err()
}

func (s *queueRepository) Dequeue(ctx context.Context, key string, count int) ([]string, error) {
	_, writer := s.clientProvider.CurrentClient(ctx)

	return writer.RPopCount(ctx, s.getKey(key), count).Result()
}

func (s *queueRepository) Len(ctx context.Context, key string) (int64, error) {
	reader, _ := s.clientProvider.CurrentClient(ctx)

	return reader.LLen(ctx, s.getKey(key)).Result()
}

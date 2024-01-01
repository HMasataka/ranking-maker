package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/HMasataka/ranking-maker/domain/repository"
	tx "github.com/HMasataka/transactor/redis"
	"github.com/redis/go-redis/v9"
)

type scoreRepository struct {
	clientProvider tx.ClientProvider
}

func NewScoreRepository(clientProvider tx.ClientProvider) repository.ScoreRepository {
	return &scoreRepository{
		clientProvider: clientProvider,
	}
}

func (s *scoreRepository) getKey(key string) string {
	return fmt.Sprintf("score:%s", key)
}

func (s *scoreRepository) Count(ctx context.Context, key string, expired time.Duration) (int64, error) {
	reader, _ := s.clientProvider.CurrentClient(ctx)

	ex := time.Now().Unix() - int64(expired.Seconds())

	return reader.ZCount(ctx, s.getKey(key), fmt.Sprintf("%d", ex), "+inf").Result()
}

func (s *scoreRepository) Increment(ctx context.Context, key, member string) error {
	_, writer := s.clientProvider.CurrentClient(ctx)

	return writer.ZAdd(ctx, s.getKey(key), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: member,
	}).Err()
}

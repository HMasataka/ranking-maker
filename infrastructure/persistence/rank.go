package persistence

import (
	"context"
	"fmt"

	"github.com/HMasataka/ranking-maker/domain/repository"
	tx "github.com/HMasataka/transactor/redis"
	"github.com/redis/go-redis/v9"
)

type rankRepository struct {
	clientProvider tx.ClientProvider
}

func NewRankRepository(clientProvider tx.ClientProvider) repository.RankRepository {
	return &rankRepository{
		clientProvider: clientProvider,
	}
}

func (s *rankRepository) getKey(key string) string {
	return fmt.Sprintf("rank:%s", key)
}

func (s *rankRepository) Add(ctx context.Context, key string, score float64, member any) error {
	_, writer := s.clientProvider.CurrentClient(ctx)

	return writer.ZAdd(ctx, s.getKey(key), redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

func (s *rankRepository) Range(ctx context.Context, key string, min, max int64) ([]redis.Z, error) {
	reader, _ := s.clientProvider.CurrentClient(ctx)

	return reader.ZRangeWithScores(ctx, s.getKey(key), min, max).Result()
}

package repository

import (
	"context"
	"time"
)

type ScoreRepository interface {
	Count(ctx context.Context, key string, expired time.Duration) (int64, error)
	Add(ctx context.Context, key string, member any) error
	RemoveOlder(ctx context.Context, key string, expired time.Duration) error
}

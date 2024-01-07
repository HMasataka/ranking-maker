package repository

import (
	"context"
)

type QueueRepository interface {
	Enqueue(ctx context.Context, key string, members ...any) error
	Dequeue(ctx context.Context, key string, count int64) ([]string, error)
	Get(ctx context.Context, key string, start, stop int64) ([]string, error)
	Len(ctx context.Context, key string) (int64, error)
}

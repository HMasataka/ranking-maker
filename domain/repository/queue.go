package repository

import (
	"context"
)

type QueueRepository interface {
	Enqueue(ctx context.Context, key string, members ...any) error
	Dequeue(ctx context.Context, key string, count int) ([]string, error)
	Len(ctx context.Context, key string) (int64, error)
}
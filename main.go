package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HMasataka/ranking-maker/application/usecase"
	"github.com/HMasataka/ranking-maker/infrastructure/persistence"
	tx "github.com/HMasataka/transactor/redis"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "localhost", 6379),
		Password: "",
		DB:       0,
	})

	return redisClient
}

func main() {
	client := NewRedisClient()
	ctx := context.Background()

	connection := tx.NewConnectionProvider(client)
	clientProvider := tx.NewClientProvider(connection)
	scoreRepository := persistence.NewScoreRepository(clientProvider)
	u := usecase.NewScoreUseCase(connection, scoreRepository)

	err := u.Increment(ctx, "key", "member1")
	if err != nil {
		panic(err)
	}

	res, err := u.Count(ctx, "key", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}

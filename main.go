package main

import (
	"context"
	"fmt"

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
	res := client.Ping(context.Background())
	fmt.Printf("%+v", res)
}

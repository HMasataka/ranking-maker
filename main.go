package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HMasataka/config"
	"github.com/HMasataka/ranking-maker/di"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewRedisConfig()
	if err != nil {
		panic(err)
	}

	scoreService := di.InitializeScoreService(cfg)
	rankService := di.InitializeRankService(cfg)

	if err := scoreService.Add(ctx, "key", "member"); err != nil {
		panic(err)
	}

	res, err := scoreService.Count(ctx, "key", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)

	if err := rankService.Add(ctx, "key", 1, "member"); err != nil {
		panic(err)
	}

	rank, err := rankService.Range(ctx, "key", 0, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", rank)
}

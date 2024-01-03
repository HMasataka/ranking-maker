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

	u := di.InitializeScoreService(cfg)

	if err := u.Add(ctx, "key", "member"); err != nil {
		panic(err)
	}

	res, err := u.Count(ctx, "key", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}

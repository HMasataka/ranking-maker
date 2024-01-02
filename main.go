package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HMasataka/ranking-maker/di"
)

func main() {
	ctx := context.Background()

	u := di.InitializeScoreService()

	err := u.Add(ctx, "key", "member1")
	if err != nil {
		panic(err)
	}

	res, err := u.Count(ctx, "key", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}

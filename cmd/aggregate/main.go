package main

import (
	"context"

	"github.com/HMasataka/config"
	"github.com/HMasataka/ranking-maker/di"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
)

type Opts struct {
	Key      string `short:"k" long:"key" description:"Redis key for save result"`
	Duration uint   `short:"d" long:"duration" description:"Aggregate duration"`
}

func ParseArgs() (*Opts, error) {
	var opts Opts
	_, err := flags.Parse(&opts)
	if err != nil {
		return nil, err
	}

	return &opts, nil
}

func main() {
	ctx := context.Background()

	cfg, err := config.NewRedisConfig()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	args, err := ParseArgs()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	usecase := di.InitializeAggregateUseCase(cfg)

	err = usecase.Execute(ctx, args.Key)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

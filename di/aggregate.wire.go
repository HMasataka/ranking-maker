//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HMasataka/config"
	"github.com/HMasataka/ranking-maker/application/usecase"
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/ranking-maker/infrastructure/persistence"
	"github.com/google/wire"
)

func InitializeAggregateUseCase(cfg *config.RedisConfig) usecase.AggregateUseCase {
	wire.Build(
		RedisClient,
		persistence.NewScoreRepository,
		persistence.NewRankRepository,
		persistence.NewQueueRepository,
		service.NewScoreService,
		service.NewRankService,
		service.NewQueueService,
		usecase.NewAggregateUseCase,
	)

	return nil
}

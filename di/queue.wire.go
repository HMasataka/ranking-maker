//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HMasataka/config"
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/ranking-maker/infrastructure/persistence"
	"github.com/google/wire"
)

func InitializeQueueService(cfg *config.RedisConfig) service.QueueService {
	wire.Build(
		RedisClient,
		persistence.NewQueueRepository,
		service.NewQueueService,
	)

	return nil
}

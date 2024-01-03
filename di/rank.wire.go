//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HMasataka/config"
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/ranking-maker/infrastructure/persistence"
	"github.com/google/wire"
)

func InitializeRankService(cfg *config.RedisConfig) service.RankService {
	wire.Build(
		RedisClient,
		persistence.NewRankRepository,
		service.NewRankService,
	)

	return nil
}

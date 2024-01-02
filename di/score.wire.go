//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/ranking-maker/infrastructure/persistence"
	"github.com/google/wire"
)

func InitializeScoreService() service.ScoreService {
	wire.Build(
		RedisClient,
		persistence.NewScoreRepository,
		service.NewScoreService,
	)

	return nil
}

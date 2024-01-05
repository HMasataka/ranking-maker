package di

import (
	"github.com/HMasataka/ranking-maker/infrastructure"
	"github.com/HMasataka/transactor/redis"
	"github.com/google/wire"
)

var RedisClient = wire.NewSet(
	infrastructure.NewRedisClient,
	redis.NewConnectionProvider,
	redis.NewClientProvider,
	redis.NewTransactor,
)

package usecase

import (
	"context"

	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/transactor"
)

type AggregateUseCase interface {
	Execute(ctx context.Context, key string) error
}

type aggregateUseCase struct {
	transactor   transactor.Transactor
	scoreService service.ScoreService
	rankService  service.RankService
	queueService service.QueueService
}

func NewAggregateUseCase(
	transactor transactor.Transactor,
	scoreService service.ScoreService,
	rankService service.RankService,
	queueService service.QueueService,
) AggregateUseCase {
	return &aggregateUseCase{
		transactor:   transactor,
		scoreService: scoreService,
		rankService:  rankService,
		queueService: queueService,
	}
}

func (c *aggregateUseCase) Execute(ctx context.Context, key string) error {
	return c.transactor.Required(ctx, func(ctx context.Context) error {
		return nil
	})
}

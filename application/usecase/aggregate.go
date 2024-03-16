package usecase

import (
	"context"
	"time"

	"github.com/HMasataka/ranking-maker/domain/entity"
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/transactor"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
)

type AggregateUseCase interface {
	Execute(ctx context.Context, key string, duration time.Duration) error
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

func (a aggregateUseCase) Execute(ctx context.Context, key string, duration time.Duration) error {
	log.Info().Time("start time", time.Now()).Send()

	if err := a.transactor.Required(ctx, func(ctx context.Context) error {
		queueLength, err := a.queueService.Len(ctx, key)
		if err != nil {
			return err
		}

		if queueLength == 0 {
			return nil
		}

		// TODO queueLength が大きすぎる場合は分割して一部読み込み
		targets, err := a.queueService.Get(ctx, key, 0, queueLength)
		if err != nil {
			return err
		}

		for i := range targets {
			var item entity.Item

			if err := json.Unmarshal([]byte(targets[i]), &item); err != nil {
				return err
			}

			score, err := a.scoreService.Count(ctx, item.ID, duration)
			if err != nil {
				return err
			}

			if err := a.rankService.Add(ctx, key, float64(score), &item); err != nil {
				return err
			}
		}

		if err := a.queueService.Delete(ctx, key, queueLength); err != nil {
			return err
		}

		if err := a.queueService.Push(ctx, key, stringToAny(targets)...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	log.Info().Time("end time", time.Now()).Send()

	return nil
}

func stringToAny(v []string) []any {
	s := make([]any, len(v))
	for i, vv := range v {
		s[i] = vv
	}
	return s
}

package usecase

import (
	"context"
	"time"

	"github.com/HMasataka/ranking-maker/domain/entity"
	"github.com/HMasataka/ranking-maker/domain/service"
	"github.com/HMasataka/transactor"
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

const limit = 100000

var jst *time.Location

func init() {
	var err error

	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func (a aggregateUseCase) Execute(ctx context.Context, key string, duration time.Duration) error {
	log.Info().Time("start time", time.Now().In(jst)).Send()

	if err := a.transactor.Required(ctx, func(ctx context.Context) error {
		queueLength, err := a.queueService.Len(ctx, key)
		if err != nil {
			return err
		}

		log.Info().Int64("queue length", queueLength).Send()

		if queueLength == 0 {
			return nil
		}

		if queueLength > limit {
			queueLength = limit
		}

		targets, err := a.queueService.Get(ctx, key, 0, queueLength)
		if err != nil {
			return err
		}

		next := make([]string, 0, len(targets))

		seen := make(map[string]struct{})

		for i := range targets {
			var item entity.Item

			if err := item.UnmarshalBinary([]byte(targets[i])); err != nil {
				return err
			}

			if _, found := seen[item.ID]; found {
				continue
			}

			seen[item.ID] = struct{}{}

			score, err := a.scoreService.Count(ctx, item.ID, duration)
			if err != nil {
				return err
			}

			log.Debug().Str("item id", item.ID).Int64("score", score).Send()

			if isToReject(score, &item) {
				if err := a.rankService.Add(ctx, key, 0, &item); err != nil {
					return err
				}

				continue
			}

			if err := a.rankService.Add(ctx, key, float64(score), &item); err != nil {
				return err
			}

			next = append(next, targets[i])
		}

		if err := a.queueService.Delete(ctx, key, queueLength); err != nil {
			return err
		}

		if err := a.queueService.Push(ctx, key, stringToAny(next)...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	log.Info().Time("end time", time.Now().In(jst)).Send()

	return nil
}

const scoreThreshold = 10

// 投稿日時が古くスコアも一定より低ければqueueから削除
func isToReject(score int64, item *entity.Item) bool {
	threshold := time.Now().In(jst).AddDate(0, 0, -7)

	if item.CreatedAt.Before(threshold) && score < scoreThreshold {
		return true
	}

	return false
}

func stringToAny(v []string) []any {
	s := make([]any, len(v))
	for i, vv := range v {
		s[i] = vv
	}
	return s
}

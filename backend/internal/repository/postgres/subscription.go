package postgres

import (
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

type subModel struct {
	Id          int
	Name        string
	Description string
	Cost        int
	Duration    int
	AccessLvl   int
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db}
}

func (s *SubscriptionRepo) GetList(ctx context.Context) ([]model.Subscription, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var subs []subModel
	var resSubs []model.Subscription
	result := s.db.WithContext(ctx).Table("dorama_set.subscription").Find(&subs)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(subs) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, s := range subs {
		tmp := model.Subscription{
			Id:          s.Id,
			Name:        s.Name,
			Description: s.Description,
			Cost:        s.Cost,
			Duration:    time.Duration(s.Duration) * constant.Day,
			AccessLvl:   s.AccessLvl,
		}
		resSubs = append(resSubs, tmp)
	}
	return resSubs, nil
}

func (s *SubscriptionRepo) GetSubscription(ctx context.Context, id int) (*model.Subscription, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetSubscription")
	defer span.End()
	var sub *subModel
	result := s.db.WithContext(ctx).Table("dorama_set.subscription").Where("id = ?", id).Take(&sub)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	res := model.Subscription{
		Id:          sub.Id,
		Name:        sub.Name,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    time.Duration(sub.Duration) * constant.Day,
		AccessLvl:   sub.AccessLvl,
	}
	return &res, nil
}

func (s *SubscriptionRepo) GetSubscriptionByPrice(ctx context.Context, price int) (*model.Subscription, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetSubscriptionByPrice")
	defer span.End()
	var sub *model.Subscription
	result := s.db.WithContext(ctx).Table("dorama_set.subscription").Where("cost = ?", price).Find(&sub)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if sub.Id == 0 {
		return nil, fmt.Errorf("db: don't exists")
	}
	res := model.Subscription{
		Id:          sub.Id,
		Name:        sub.Name,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    sub.Duration * constant.Day,
		AccessLvl:   sub.AccessLvl,
	}

	return &res, nil
}

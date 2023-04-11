package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

type subModel struct {
	Id          int
	Description string
	Cost        int
	Duration    int
}

func (s SubscriptionRepo) GetList() ([]model.Subscription, error) {
	var subs []subModel
	var resSubs []model.Subscription
	result := s.db.Table("dorama_set.subscription").Find(&subs)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, s := range subs {
		tmp := model.Subscription{
			Id:          s.Id,
			Description: s.Description,
			Cost:        s.Cost,
			Duration:    time.Duration(s.Duration) * time.Second,
		}
		resSubs = append(resSubs, tmp)
	}
	return resSubs, nil
}

func (s SubscriptionRepo) GetSubscription(id int) (*model.Subscription, error) {
	var sub *subModel
	result := s.db.Table("dorama_set.subscription").Where("id = ?", id).Find(&sub)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	// todo вынести ошибку
	if sub.Id == 0 {
		return nil, fmt.Errorf("db: don't exists")
	}
	res := model.Subscription{
		Id:          sub.Id,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    time.Duration(sub.Duration) * time.Second}
	return &res, nil
}

func (s SubscriptionRepo) GetSubscriptionByPrice(price int) (*model.Subscription, error) {
	var sub *model.Subscription
	result := s.db.Table("dorama_set.subscription").Where("cost = ?", price).Find(&sub)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if sub.Id == 0 {
		return nil, fmt.Errorf("db: don't exists")
	}
	res := model.Subscription{
		Id:          sub.Id,
		Description: sub.Description,
		Cost:        sub.Cost,
		Duration:    time.Duration(sub.Duration) * time.Second}

	return &res, nil
}

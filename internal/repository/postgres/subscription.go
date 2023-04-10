package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type SubscritionRepo struct {
	db *gorm.DB
}

func (SubscritionRepo) GetList() ([]model.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (SubscritionRepo) GetSubscription(id int) (*model.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

func (SubscritionRepo) GetSubscriptionByPrice(price int) (*model.Subscription, error) {
	//TODO implement me
	panic("implement me")
}

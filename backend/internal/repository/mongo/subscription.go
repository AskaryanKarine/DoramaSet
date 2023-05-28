package mongo

import (
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriptionRepo struct {
	db *mongo.Database
}

func NewSubscriptionRepo(db *mongo.Database) *SubscriptionRepo {
	return &SubscriptionRepo{db}
}

func (SubscriptionRepo) GetList() ([]model.Subscription, error) {
	// TODO implement me
	panic("implement me")
}

func (SubscriptionRepo) GetSubscription(id int) (*model.Subscription, error) {
	// TODO implement me
	panic("implement me")
}

func (SubscriptionRepo) GetSubscriptionByPrice(price int) (*model.Subscription, error) {
	// TODO implement me
	panic("implement me")
}

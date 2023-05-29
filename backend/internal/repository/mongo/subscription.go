package mongo

import (
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SubscriptionRepo struct {
	db *mongo.Database
}

type subModel struct {
	Id          int    `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Cost        int    `bson:"cost"`
	Duration    int    `bson:"duration"`
	AccessLvl   int    `bson:"access_lvl"`
}

func NewSubscriptionRepo(db *mongo.Database) *SubscriptionRepo {
	return &SubscriptionRepo{db}
}

func getSubLogicModel(m subModel) *model.Subscription {
	return &model.Subscription{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
		Cost:        m.Cost,
		Duration:    time.Duration(m.Duration) * constant.Day,
		AccessLvl:   m.AccessLvl,
	}
}

func (s *SubscriptionRepo) GetList() ([]model.Subscription, error) {
	var (
		resDB []subModel
		res   []model.Subscription
	)
	collection := s.db.Collection("subscription")
	opts := options.Find().SetSort(bson.D{{"id", 1}})
	cur, err := collection.Find(nil, bson.D{{}}, opts)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		res = append(res, *getSubLogicModel(r))
	}
	return res, nil
}

func (s *SubscriptionRepo) GetSubscription(id int) (*model.Subscription, error) {
	var resDB subModel

	collection := s.db.Collection("subscription")
	filter := bson.D{{"id", id}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	return getSubLogicModel(resDB), nil
}

func (s *SubscriptionRepo) GetSubscriptionByPrice(price int) (*model.Subscription, error) {
	var resDB subModel

	collection := s.db.Collection("subscription")
	filter := bson.D{{"cost", price}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	return getSubLogicModel(resDB), nil
}

package mongo

import (
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepo struct {
	db *mongo.Database
}

func NewReviewRepo(db *mongo.Database) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (ReviewRepo) GetAllReview(idD int) ([]model.Review, error) {
	// TODO implement me
	panic("implement me")
}

func (ReviewRepo) CreateReview(idD int, record *model.Review) error {
	// TODO implement me
	panic("implement me")
}

func (ReviewRepo) DeleteReview(username string, idD int) error {
	// TODO implement me
	panic("implement me")
}

func (ReviewRepo) AggregateRate(idD int) (float64, int, error) {
	// TODO implement me
	panic("implement me")
}

func (ReviewRepo) GetReview(username string, idD int) (*model.Review, error) {
	// TODO implement me
	panic("implement me")
}

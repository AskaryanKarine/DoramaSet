package mongo

import (
	errors2 "DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepo struct {
	db *mongo.Database
}

type reviewModel struct {
	Username string `bson:"username"`
	Mark     int    `bson:"mark"`
	Content  string `bson:"content"`
}

func NewReviewRepo(db *mongo.Database) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) GetAllReview(ctx context.Context, idD int) ([]model.Review, error) {
	var res []model.Review
	collection := r.db.Collection("dorama")
	filter := bson.D{{"id", idD}}
	var result struct {
		Review []reviewModel `bson:"review"`
	}
	err := collection.FindOne(nil, filter).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, d := range result.Review {
		m := model.Review{
			Username: d.Username,
			Mark:     d.Mark,
			Content:  d.Content,
		}
		res = append(res, m)
	}
	return res, nil
}

func (r *ReviewRepo) CreateReview(ctx context.Context, idD int, record *model.Review) error {
	var resDB struct {
		Review []reviewModel `bson:"review"`
	}
	collection := r.db.Collection("dorama")
	helpFilter := bson.D{{"id", idD}, {"review.username", record.Username}}
	err := collection.FindOne(nil, helpFilter).Decode(&resDB)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	if len(resDB.Review) > 0 {
		return errors2.ErrorExistInDB
	}
	newReview := reviewModel{
		Username: record.Username,
		Mark:     record.Mark,
		Content:  record.Content,
	}
	filter := bson.D{{"id", idD}}
	update := bson.D{{"$push", bson.D{{"review", newReview}}}}
	_, err = collection.UpdateOne(nil, filter, update)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (r *ReviewRepo) DeleteReview(ctx context.Context, username string, idD int) error {
	collection := r.db.Collection("dorama")
	filter := bson.D{{"id", idD}}
	update := bson.D{{"$pull", bson.D{{"review", bson.D{{"username", username}}}}}}
	_, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (r *ReviewRepo) AggregateRate(ctx context.Context, idD int) (float64, int, error) {
	collection := r.db.Collection("dorama")
	projectStage := bson.D{{"$project",
		bson.D{{"id", "$id"},
			{"cnt",
				bson.D{
					{"$cond", bson.A{
						bson.D{{"$isArray", "$review"}},
						bson.D{{"$size", "$review"}},
						0}}},
			},
			{"rate",
				bson.D{
					{"$cond", bson.A{
						bson.D{{"$isArray", "$review"}},
						bson.D{{"$avg", "$review.mark"}},
						0}}},
			},
		},
	}}
	matchStage := bson.D{{"$match", bson.D{{"id", idD}}}}
	cur, err := collection.Aggregate(nil, mongo.Pipeline{projectStage, matchStage})
	if err != nil {
		return 0, 0, fmt.Errorf("db: %w", err)
	}
	var resDB []struct {
		Cnt  int     `bson:"cnt"`
		Rate float64 `bson:"rate"`
	}
	if err = cur.All(nil, &resDB); err != nil {
		return 0, 0, fmt.Errorf("db: %w", err)
	}

	return resDB[0].Rate, resDB[0].Cnt, nil
}

func (r *ReviewRepo) GetReview(ctx context.Context, username string, idD int) (*model.Review, error) {
	var resDB struct {
		Review reviewModel `bson:"review"`
	}
	collection := r.db.Collection("dorama")
	helpFilter := bson.D{{"id", idD}, {"review.username", username}}
	err := collection.FindOne(nil, helpFilter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	review := model.Review{
		Username: resDB.Review.Username,
		Mark:     resDB.Review.Mark,
		Content:  resDB.Review.Content,
	}
	return &review, nil
}

package mongo

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/interfaces/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func Open(cfg *config.Config) (*repository.AllRepository, error) {
	dsn := "mongodb://%s:%s@%s:%d"
	dsn = fmt.Sprintf(dsn, cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, fmt.Errorf("open mongo client: %w", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect mongo client: %w", err)
	}
	// defer func() {
	// 	_ = client.Disconnect(ctx)
	// }()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	all := create(client.Database("DoramaSet"))
	return all, nil
}

func create(db *mongo.Database) *repository.AllRepository {
	picRepo := NewPictureRepo(db)
	eRepo := NewEpisodeRepo(db)
	revRepo := NewReviewRepo(db)
	dRepo := NewDoramaRepo(db, picRepo, eRepo, revRepo)
	lRepo := NewListRepo(db, dRepo)
	staffRepo := NewStaffRepo(db, picRepo)
	subRepo := NewSubscriptionRepo(db)
	uRepo := NewUserRepo(db, subRepo, lRepo)

	all := repository.AllRepository{
		Dorama:       dRepo,
		Review:       revRepo,
		Episode:      eRepo,
		List:         lRepo,
		Picture:      picRepo,
		Subscription: subRepo,
		Staff:        staffRepo,
		User:         uRepo,
	}
	return &all
}

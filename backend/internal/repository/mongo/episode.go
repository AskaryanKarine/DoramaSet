package mongo

import (
	errors2 "DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EpisodeRepo struct {
	db *mongo.Database
}

func NewEpisodeRepo(db *mongo.Database) *EpisodeRepo {
	return &EpisodeRepo{db}
}

type episodeModel struct {
	ID         int `bson:"id"`
	IDDorama   int `bson:"id_dorama"`
	NumSeason  int `bson:"num_season"`
	NumEpisode int `bson:"num_episode"`
}

func getEpisodeLogicModel(m episodeModel) *model.Episode {
	return &model.Episode{
		Id:         m.ID,
		NumSeason:  m.NumSeason,
		NumEpisode: m.NumEpisode,
	}
}

func (e *EpisodeRepo) GetList(idDorama int) ([]model.Episode, error) {
	var (
		resDB []episodeModel
		res   []model.Episode
	)

	collection := e.db.Collection("_episode")
	filter := bson.D{{"id_dorama", idDorama}}
	opts := options.Find().SetSort(bson.D{{"id", 1}})

	cur, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db_find: %w", err)
	}

	if err := cur.All(context.TODO(), &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	for _, r := range resDB {
		tmp := model.Episode{
			Id:         r.ID,
			NumSeason:  r.NumSeason,
			NumEpisode: r.NumEpisode,
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (e *EpisodeRepo) GetWatchingList(username string, idD int) ([]model.Episode, error) {
	// TODO implement me
	panic("implement me")
}

func (e *EpisodeRepo) GetEpisode(id int) (*model.Episode, error) {
	var resDB episodeModel

	collection := e.db.Collection("_episode")
	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	return getEpisodeLogicModel(resDB), nil
}

func (e *EpisodeRepo) MarkEpisode(idEp int, username string) error {
	type query struct {
		Username string `bson:"username"`
		Episode  int    `bson:"episode"`
	}
	var resDB query

	collection := e.db.Collection("_user_watched_episode")
	helpFilter := bson.D{{"username", username}, {"episode", idEp}}
	err := collection.FindOne(nil, helpFilter).Decode(&resDB)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	if err == nil {
		return errors2.ErrorExistInDB
	}
	_, err = collection.InsertOne(nil, query{
		Username: username,
		Episode:  idEp,
	})
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (e *EpisodeRepo) CreateEpisode(episode model.Episode, idD int) (int, error) {
	var (
		maxIDEp episodeModel
	)
	collection := e.db.Collection("_episode")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	err := collection.FindOne(context.TODO(), filter, opts).Decode(&maxIDEp)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	newEpisode := episodeModel{
		ID:         maxIDEp.ID + 1,
		IDDorama:   idD,
		NumSeason:  episode.NumSeason,
		NumEpisode: episode.NumEpisode,
	}
	_, err = collection.InsertOne(context.TODO(), newEpisode)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	return newEpisode.ID, nil
}

func (e *EpisodeRepo) DeleteEpisode(id int) error {
	collection := e.db.Collection("_episode")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

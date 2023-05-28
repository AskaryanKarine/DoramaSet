package mongo

import (
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type EpisodeRepo struct {
	db *mongo.Database
}

func NewEpisodeRepo(db *mongo.Database) *EpisodeRepo {
	return &EpisodeRepo{db}
}

func (EpisodeRepo) GetList(idDorama int) ([]model.Episode, error) {
	// TODO implement me
	panic("implement me")
}

func (EpisodeRepo) GetWatchingList(username string, idD int) ([]model.Episode, error) {
	// TODO implement me
	panic("implement me")
}

func (EpisodeRepo) GetEpisode(id int) (*model.Episode, error) {
	// TODO implement me
	panic("implement me")
}

func (EpisodeRepo) MarkEpisode(idEp int, username string) error {
	// TODO implement me
	panic("implement me")
}

func (EpisodeRepo) CreateEpisode(episode model.Episode, idD int) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (EpisodeRepo) DeleteEpisode(id int) error {
	// TODO implement me
	panic("implement me")
}

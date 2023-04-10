package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type EpisodeRepo struct {
	db *gorm.DB
}

func (EpisodeRepo) GetList(idDorama int) ([]model.Episode, error) {
	//TODO implement me
	panic("implement me")
}

func (EpisodeRepo) GetEpisode(id int) (*model.Episode, error) {
	//TODO implement me
	panic("implement me")
}

func (EpisodeRepo) MarkEpisode(idEp int, username string) error {
	//TODO implement me
	panic("implement me")
}

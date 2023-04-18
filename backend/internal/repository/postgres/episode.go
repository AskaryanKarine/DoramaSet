package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type EpisodeRepo struct {
	db *gorm.DB
}

type episodeModel struct {
	ID         int
	IdDorama   int
	NumSeason  int
	NumEpisode int
}

type markEpisode struct {
	Username  string
	IdEpisode int
}

func NewEpisodeRepo(db *gorm.DB) *EpisodeRepo {
	return &EpisodeRepo{db}
}

func (e *EpisodeRepo) GetList(idDorama int) ([]model.Episode, error) {
	var res []model.Episode
	result := e.db.Table("dorama_set.episode").Where("id_dorama = ?", idDorama).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

func (e *EpisodeRepo) GetEpisode(id int) (*model.Episode, error) {
	var res *model.Episode
	result := e.db.Table("dorama_set.episode").Where("id = ?", id).Take(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	return res, nil
}

func (e *EpisodeRepo) CreateEpisode(episode model.Episode, idD int) (int, error) {
	m := episodeModel{
		IdDorama:   idD,
		NumSeason:  episode.NumSeason,
		NumEpisode: episode.NumEpisode,
	}
	result := e.db.Table("dorama_set.episode").
		Omit("id").
		Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (e *EpisodeRepo) DeleteEpisode(id int) error {
	result := e.db.Table("dorama_set.episode").Where("id = ?", id).Delete(&model.Episode{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (e *EpisodeRepo) MarkEpisode(idEp int, username string) error {
	m := markEpisode{Username: username, IdEpisode: idEp}
	result := e.db.Table("dorama_set.userepisode").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

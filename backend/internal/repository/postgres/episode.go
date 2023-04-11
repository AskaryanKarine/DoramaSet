package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type EpisodeRepo struct {
	db *gorm.DB
}

func (e EpisodeRepo) GetList(idDorama int) ([]model.Episode, error) {
	var res []model.Episode
	result := e.db.Table("dorama_set.episode").Where("id_dorama = ?", idDorama).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	// todo new error
	if len(res) == 0 {
		return nil, fmt.Errorf("db: dont exists")
	}
	return res, nil
}

func (e EpisodeRepo) GetEpisode(id int) (*model.Episode, error) {
	var res *model.Episode
	result := e.db.Table("dorama_set.episode").Where("id = ?", id).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	// todo new error
	if res.Id < 1 {
		return nil, fmt.Errorf("db: don't exists")
	}
	return res, nil
}

type episodeModel struct {
	ID         int
	IdDorama   int
	NumSeason  int
	NumEpisode int
}

func (e EpisodeRepo) CreateEpisode(episode model.Episode, idD int) (int, error) {
	m := episodeModel{
		IdDorama:   idD,
		NumSeason:  episode.NumSeason,
		NumEpisode: episode.NumEpisode,
	}
	result := e.db.Table("dorama_set.episode").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (e EpisodeRepo) DeleteEpisode(episode model.Episode) error {
	result := e.db.Table("dorama_set.episode").Where("id = ?", episode.Id).Delete(&model.Episode{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

type markEpisode struct {
	Username  string
	IdEpisode int
}

func (e EpisodeRepo) MarkEpisode(idEp int, username string) error {
	m := markEpisode{Username: username, IdEpisode: idEp}
	result := e.db.Table("dorama_set.userepisode").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

//func (e EpisodeRepo) DeleteMarkEpisode(idEp int, username string) error {
//	result := e.db.Table("dorama_set.userepisode").Where("username = ? and id_episode = ?", username, idEp).Delete()
//}

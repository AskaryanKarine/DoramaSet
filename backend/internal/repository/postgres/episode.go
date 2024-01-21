package postgres

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
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

func (e *EpisodeRepo) GetList(ctx context.Context, idDorama int) ([]model.Episode, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var res []model.Episode
	result := e.db.WithContext(ctx).Table("dorama_set.episode").Where("id_dorama = ?", idDorama).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

func (e *EpisodeRepo) GetEpisode(ctx context.Context, id int) (*model.Episode, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var res *model.Episode
	result := e.db.WithContext(ctx).Table("dorama_set.episode").Where("id = ?", id).Take(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	return res, nil
}

func (e *EpisodeRepo) CreateEpisode(ctx context.Context, episode model.Episode, idD int) (int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	m := episodeModel{
		IdDorama:   idD,
		NumSeason:  episode.NumSeason,
		NumEpisode: episode.NumEpisode,
	}
	result := e.db.WithContext(ctx).Table("dorama_set.episode").
		Omit("id").
		Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (e *EpisodeRepo) DeleteEpisode(ctx context.Context, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	result := e.db.WithContext(ctx).Table("dorama_set.episode").Where("id = ?", id).Delete(&model.Episode{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (e *EpisodeRepo) MarkEpisode(ctx context.Context, idEp int, username string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	m := markEpisode{Username: username, IdEpisode: idEp}
	result := e.db.WithContext(ctx).Table("dorama_set.userepisode").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (e *EpisodeRepo) GetWatchingList(ctx context.Context, username string, idD int) ([]model.Episode, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var resDB []model.Episode
	result := e.db.WithContext(ctx).Table("dorama_set.episode e").Select("e.*").
		Joins("join dorama_set.userepisode ue on ue.id_episode = e.id").
		Where("ue.username = ? and e.id_dorama = ?", username, idD).Find(&resDB)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	return resDB, nil
}

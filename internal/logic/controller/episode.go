package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/interfaces"
)

type EpisodeController struct {
	repo  interfaces.IEpisodeRepo
	urepo interfaces.IUserRepo
}

func (e *EpisodeController) GetEpisodeList(idD int) ([]model.Episode, error) {
	return e.repo.GetList(idD)
}

func (e *EpisodeController) GetEpisode(id int) (model.Episode, error) {
	return e.repo.GetEpisode(id)
}

func (e *EpisodeController) MarkWatchingEpisode(idEp int, username string) error {
	_, err := e.urepo.GetUser(username)
	if err != nil {
		return err
	}
	return e.repo.MarkEpisode(idEp, username)
}

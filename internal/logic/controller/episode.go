package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
)

type EpisodeController struct {
	repo  interfaces.IEpisodeRepo
	urepo interfaces.IUserRepo
}

func (e *EpisodeController) GetEpisodeList(idD int) ([]model.Episode, error) {
	res, err := e.repo.GetList(idD)
	return res, err
}

func (e *EpisodeController) GetEpisode(id int) (model.Episode, error) {
	res, err := e.repo.GetEpisode(id)
	return res, err
}

func (e *EpisodeController) MarkWatchingEpisode(idEp int, username string) error {
	_, err := e.urepo.GetUser(username)
	if err != nil {
		return err
	}
	return e.repo.MarkEpisode(idEp, username)
}

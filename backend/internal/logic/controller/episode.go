package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type EpisodeController struct {
	repo repository.IEpisodeRepo
	uc   controller.IUserController
}

func NewEpisodeController(ERepo repository.IEpisodeRepo, uc controller.IUserController) *EpisodeController {
	return &EpisodeController{
		repo: ERepo,
		uc:   uc,
	}
}

func (e *EpisodeController) GetEpisodeList(idD int) ([]model.Episode, error) {
	res, err := e.repo.GetList(idD)
	if err != nil {
		return nil, fmt.Errorf("getList: %w", err)
	}
	return res, nil
}

func (e *EpisodeController) GetEpisode(id int) (*model.Episode, error) {
	res, err := e.repo.GetEpisode(id)
	if err != nil {
		return nil, fmt.Errorf("getEpisode: %w", err)
	}
	return res, nil
}

func (e *EpisodeController) MarkWatchingEpisode(token string, idEp int) error {
	user, err := e.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}
	err = e.repo.MarkEpisode(idEp, user.Username)
	if err != nil {
		return fmt.Errorf("markEpisode: %w", err)
	}
	return nil
}

func (e *EpisodeController) CreateEpisode(record model.Episode, idD int) error {
	_, err := e.repo.CreateEpisode(record, idD)
	if err != nil {
		return fmt.Errorf("createEpisode: %w", err)
	}
	return nil
}

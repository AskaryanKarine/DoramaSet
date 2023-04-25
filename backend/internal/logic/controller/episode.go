package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/sirupsen/logrus"
)

type EpisodeController struct {
	repo repository.IEpisodeRepo
	uc   controller.IUserController
	log  *logrus.Logger
}

func NewEpisodeController(ERepo repository.IEpisodeRepo, uc controller.IUserController,
	log *logrus.Logger) *EpisodeController {
	return &EpisodeController{
		repo: ERepo,
		uc:   uc,
		log:  log,
	}
}

func (e *EpisodeController) GetEpisodeList(idD int) ([]model.Episode, error) {
	res, err := e.repo.GetList(idD)
	if err != nil {
		e.log.Warnf("get episode list, get list err: %s, value %d", err, idD)
		return nil, fmt.Errorf("getList: %w", err)
	}
	e.log.Infof("got episode list, value: %d", idD)
	return res, nil
}

func (e *EpisodeController) GetEpisode(id int) (*model.Episode, error) {
	res, err := e.repo.GetEpisode(id)
	if err != nil {
		e.log.Warnf("get episode, get err: %s, value %d", err, id)
		return nil, fmt.Errorf("getEpisode: %w", err)
	}
	e.log.Infof("got episode, value: %d", id)
	return res, nil
}

func (e *EpisodeController) MarkWatchingEpisode(token string, idEp int) error {
	user, err := e.uc.AuthByToken(token)
	if err != nil {
		e.log.Warnf("mark wath ep, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	err = e.repo.MarkEpisode(idEp, user.Username)
	if err != nil {
		e.log.Warnf("mark wath ep, mark err: %s, username %s, value %d", err, user.Username, idEp)
		return fmt.Errorf("markEpisode: %w", err)
	}
	e.log.Infof("marked watch episode, username %s, value %d", user.Username, idEp)
	return nil
}

func (e *EpisodeController) CreateEpisode(record model.Episode, idD int) error {
	_, err := e.repo.CreateEpisode(record, idD)
	if err != nil {
		e.log.Warnf("create episode err %s, value %v %d", err, record, idD)
		return fmt.Errorf("createEpisode: %w", err)
	}
	e.log.Infof("created episode value %v %d", record, idD)
	return nil
}

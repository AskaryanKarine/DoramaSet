package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
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

func (e *EpisodeController) GetEpisodeList(ctx context.Context, idD int) ([]model.Episode, error) {
	ctxLog, spanLog := tracing.StartSpanFromContext(ctx, "LOG getEpisodeList")
	defer spanLog.End()
	ctx, span := tracing.StartSpanFromContext(ctx, "BL getEpisodeList")
	defer span.End()
	res, err := e.repo.GetList(ctx, idD)
	if err != nil {
		e.log.WithContext(ctxLog).Warnf("get episode list, get list err: %s, value %d", err, idD)
		return nil, fmt.Errorf("getList: %w", err)
	}
	e.log.WithContext(ctxLog).Infof("got episode list, value: %d", idD)
	return res, nil
}

func (e *EpisodeController) GetEpisode(ctx context.Context, id int) (*model.Episode, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL getEpisode")
	defer span.End()
	res, err := e.repo.GetEpisode(ctx, id)
	ctx, spanLog := tracing.StartSpanFromContext(ctx, "LOG getEpisodeList")
	defer spanLog.End()
	if err != nil {
		e.log.WithContext(ctx).Warnf("get episode, get err: %s, value %d", err, id)
		return nil, fmt.Errorf("getEpisode: %w", err)
	}
	e.log.WithContext(ctx).Infof("got episode, value: %d", id)
	return res, nil
}

func (e *EpisodeController) MarkWatchingEpisode(ctx context.Context, token string, idEp int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL markWatchingEpisode")
	defer span.End()
	user, err := e.uc.AuthByToken(ctx, token)
	if err != nil {
		e.log.Warnf("mark wath ep, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	err = e.repo.MarkEpisode(ctx, idEp, user.Username)
	if err != nil {
		e.log.Warnf("mark wath ep, mark err: %s, username %s, value %d", err, user.Username, idEp)
		return fmt.Errorf("markEpisode: %w", err)
	}
	e.log.Infof("marked watch episode, username %s, value %d", user.Username, idEp)
	return nil
}

func (e *EpisodeController) CreateEpisode(ctx context.Context, token string, record *model.Episode, idD int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL createEpisode")
	defer span.End()
	user, err := e.uc.AuthByToken(ctx, token)
	if err != nil {
		e.log.Warnf("update dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !user.IsAdmin {
		e.log.Warnf("update dorama, access err: %s username %s", err, user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}
	id, err := e.repo.CreateEpisode(ctx, *record, idD)
	if err != nil {
		e.log.Warnf("create episode err %s, value %v %d", err, record, idD)
		return fmt.Errorf("createEpisode: %w", err)
	}
	record.Id = id
	e.log.Infof("created episode value %v %d", record, idD)
	return nil
}

func (e *EpisodeController) GetWatchingEpisode(ctx context.Context, token string, idD int) ([]model.Episode, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL getWatchingEpisode")
	defer span.End()
	user, err := e.uc.AuthByToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("authByToken: %w", err)
	}

	res, err := e.repo.GetWatchingList(ctx, user.Username, idD)
	if err != nil {
		return nil, fmt.Errorf("getWatchingList: %w", err)
	}

	return res, nil
}

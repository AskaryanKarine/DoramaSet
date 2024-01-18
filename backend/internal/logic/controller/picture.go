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

type PictureController struct {
	repo repository.IPictureRepo
	uc   controller.IUserController
	log  *logrus.Logger
}

func NewPictureController(PRepo repository.IPictureRepo, uc controller.IUserController,
	log *logrus.Logger) *PictureController {
	return &PictureController{
		repo: PRepo,
		uc:   uc,
		log:  log,
	}
}

func (p *PictureController) GetListByDorama(ctx context.Context, idD int) ([]model.Picture, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetListByDorama")
	defer span.End()
	res, err := p.repo.GetListDorama(ctx, idD)
	if err != nil {
		p.log.Warnf("get pic list by dorama err %s, value %d", err, idD)
		return nil, fmt.Errorf("getByDorama: %w", err)
	}
	p.log.Infof("got list pic by dorama valye %d", idD)
	return res, nil
}

func (p *PictureController) GetListByStaff(ctx context.Context, idS int) ([]model.Picture, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetListByStaff")
	defer span.End()
	res, err := p.repo.GetListStaff(ctx, idS)
	if err != nil {
		p.log.Warnf("get pic list by staff err %s, value %d", err, idS)
		return nil, fmt.Errorf("getByStaff: %w", err)
	}
	p.log.Infof("get list pic by staff value %d", idS)
	return res, nil
}

func (p *PictureController) CreatePicture(ctx context.Context, token string, record *model.Picture) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL CreatePicture")
	defer span.End()
	user, err := p.uc.AuthByToken(ctx, token)
	if err != nil {
		p.log.Warnf("create picture auth err %s, token %s, value %v", err, token, record)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		p.log.Warnf("create picture access err, user %s, value %v", user.Username, record)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	id, err := p.repo.CreatePicture(ctx, *record)
	record.Id = id
	if err != nil {
		p.log.Warnf("create picture err %s, value %v", err, record)
		return fmt.Errorf("createPicture: %w", err)
	}
	p.log.Infof("create picture value %v", record)
	return nil
}

func (p *PictureController) AddPictureToStaff(ctx context.Context, token string, record model.Picture, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL AddPictureToStaff")
	defer span.End()
	user, err := p.uc.AuthByToken(ctx, token)
	if err != nil {
		p.log.Warnf("add picture to staff auth err %s, token %s, value %v, %d", err, token, record, id)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		p.log.Warnf("add picture to staff access err, user %s, value %v, %d", user.Username, record, id)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = p.repo.AddPictureToStaff(ctx, record, id)
	if err != nil {
		p.log.Warnf("add picture to staff err %s, user %s, value %v, %d", err, user.Username, record, id)
		return fmt.Errorf("addPictureToStaff: %w", err)
	}
	p.log.Infof("added picture to staff username %s, value %v, %d", user.Username, record, id)
	return nil
}
func (p *PictureController) AddPictureToDorama(ctx context.Context, token string, record model.Picture, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL AddPictureToDorama")
	defer span.End()
	user, err := p.uc.AuthByToken(ctx, token)
	if err != nil {
		p.log.Warnf("add picture to dorama auth err %s, token %s, value %v, %d", err, token, record, id)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		p.log.Warnf("add picture to dorama access err, user %s, value %v, %d", user.Username, record, id)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = p.repo.AddPictureToDorama(ctx, record, id)
	if err != nil {
		p.log.Warnf("add picture to dorama err %s, user %s, value %v, %d", err, user.Username, record, id)
		return fmt.Errorf("createPicture: %w", err)
	}
	p.log.Infof("add picture to dorama user %s, value %v %d", user.Username, record, id)
	return nil
}

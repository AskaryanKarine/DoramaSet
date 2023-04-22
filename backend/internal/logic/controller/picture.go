package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type PictureController struct {
	repo repository.IPictureRepo
	uc   controller.IUserController
}

func NewPictureController(PRepo repository.IPictureRepo, uc controller.IUserController) *PictureController {
	return &PictureController{
		repo: PRepo,
		uc:   uc,
	}
}

func (p *PictureController) GetListByDorama(idD int) ([]model.Picture, error) {
	res, err := p.repo.GetListDorama(idD)
	if err != nil {
		return nil, fmt.Errorf("getByDorama: %w", err)
	}
	return res, nil
}

func (p *PictureController) GetListByStaff(idS int) ([]model.Picture, error) {
	res, err := p.repo.GetListStaff(idS)
	if err != nil {
		return nil, fmt.Errorf("getByStaff: %w", err)
	}
	return res, nil
}

func (p *PictureController) CreatePicture(token string, record model.Picture) error {
	user, err := p.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	_, err = p.repo.CreatePicture(record)
	if err != nil {
		return fmt.Errorf("createPicture: %w", err)
	}
	return nil
}

func (p *PictureController) AddPictureToStaff(token string, record model.Picture, id int) error {
	user, err := p.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = p.repo.AddPictureToStaff(record, id)
	if err != nil {
		return fmt.Errorf("addPictureToStaff: %w", err)
	}
	return nil
}
func (p *PictureController) AddPictureToDorama(token string, record model.Picture, id int) error {
	user, err := p.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = p.repo.AddPictureToDorama(record, id)
	if err != nil {
		return fmt.Errorf("createPicture: %w", err)
	}
	return nil
}

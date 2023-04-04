package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
	"fmt"
)

type PictureController struct {
	repo interfaces.IPictureRepo
	uc   interfaces.IUserController
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
		return err
	}

	if !user.IsAdmin {
		return errors.New("createPic: low level of access")
	}

	err = p.repo.CreatePicture(record)
	if err != nil {
		return fmt.Errorf("createPic: %w", err)
	}
	return nil
}

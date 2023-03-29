package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
)

type PictureController struct {
	repo  interfaces.IPictureRepo
	urepo interfaces.IUserRepo
}

func (p *PictureController) GetListByDorama(idD int) ([]model.Picture, error) {
	res, err := p.repo.GetListDorama(idD)
	return res, err
}

func (p *PictureController) GetListByStaff(idS int) ([]model.Picture, error) {
	res, err := p.repo.GetListStaff(idS)
	return res, err
}

func (p *PictureController) CreatePicture(username string, record model.Picture) error {
	user, err := p.urepo.GetUser(username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return errors.New("low level of access")
	}

	return p.repo.CreatePicture(record)
}

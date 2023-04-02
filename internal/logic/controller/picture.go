package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/interfaces"
	"errors"
)

type PictureController struct {
	repo  interfaces.IPictureRepo
	urepo interfaces.IUserRepo
}

func (p *PictureController) GetListByDorama(idD int) ([]model.Picture, error) {
	return p.repo.GetListDorama(idD)
}

func (p *PictureController) GetListByStaff(idS int) ([]model.Picture, error) {
	return p.repo.GetListStaff(idS)
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

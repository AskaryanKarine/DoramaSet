package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
)

type DoramaController struct {
	repo  interfaces.IDoramaRepo
	urepo interfaces.IUserRepo
}

func (d *DoramaController) GetAll() ([]model.Dorama, error) {
	res, err := d.repo.GetList()
	return res, err
}

func (d *DoramaController) GetByName(name string) ([]model.Dorama, error) {
	res, err := d.repo.GetListName(name)
	return res, err
}

func (d *DoramaController) GetById(id int) (model.Dorama, error) {
	res, err := d.repo.GetDorama(id)
	return res, err
}

func (d *DoramaController) CreateDorama(username string, record model.Dorama) error {
	user, err := d.urepo.GetUser(username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return errors.New("low level os access")
	}

	return d.repo.CreateDorama(record)
}

func (d *DoramaController) UpdateDorama(username string, record model.Dorama) error {
	user, err := d.urepo.GetUser(username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return errors.New("low level os access")
	}

	return d.repo.UpdateDorama(record)
}

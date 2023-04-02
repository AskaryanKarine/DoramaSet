package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/interfaces"
	"errors"
)

type DoramaController struct {
	repo  interfaces.IDoramaRepo
	urepo interfaces.IUserRepo
}

func (d *DoramaController) GetAll() ([]model.Dorama, error) {
	return d.repo.GetList()
}

func (d *DoramaController) GetByName(name string) ([]model.Dorama, error) {
	return d.repo.GetListName(name)
}

func (d *DoramaController) GetById(id int) (model.Dorama, error) {
	return d.repo.GetDorama(id)
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

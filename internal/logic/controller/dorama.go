package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
	"fmt"
)

type DoramaController struct {
	repo interfaces.IDoramaRepo
	uc   interfaces.IUserController
}

func (d *DoramaController) GetAll() ([]model.Dorama, error) {
	res, err := d.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("getDorama: %w", err)
	}
	return res, nil
}

func (d *DoramaController) GetByName(name string) ([]model.Dorama, error) {
	res, err := d.repo.GetListName(name)
	if err != nil {
		return nil, fmt.Errorf("getByName: %w", err)
	}
	return res, nil
}

func (d *DoramaController) GetById(id int) (*model.Dorama, error) {
	res, err := d.repo.GetDorama(id)
	if err != nil {
		return nil, fmt.Errorf("getById: %w", err)
	}
	return res, nil
}

func (d *DoramaController) CreateDorama(token string, record model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return errors.New("createDorama: low level os access")
	}

	err = d.repo.CreateDorama(record)
	return fmt.Errorf("createDorama: %w", err)
}

func (d *DoramaController) UpdateDorama(token string, record model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return errors.New("updateDorama: low level os access")
	}

	err = d.repo.UpdateDorama(record)
	return fmt.Errorf("updateDorama: %w", err)
}

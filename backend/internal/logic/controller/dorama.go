package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type DoramaController struct {
	repo repository.IDoramaRepo
	uc   controller.IUserController
}

func (d *DoramaController) GetAll() ([]model.Dorama, error) {
	res, err := d.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("getList: %w", err)
	}
	return res, nil
}

func (d *DoramaController) GetByName(name string) ([]model.Dorama, error) {
	res, err := d.repo.GetListName(name)
	if err != nil {
		return nil, fmt.Errorf("getListName: %w", err)
	}
	return res, nil
}

func (d *DoramaController) GetById(id int) (*model.Dorama, error) {
	res, err := d.repo.GetDorama(id)
	if err != nil {
		return nil, fmt.Errorf("getDorama: %w", err)
	}
	return res, nil
}

func (d *DoramaController) CreateDorama(token string, record model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}
	// TODO +adminAccessError
	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = d.repo.CreateDorama(record)
	if err != nil {
		return fmt.Errorf("createDorama: %w", err)
	}
	return nil
}

func (d *DoramaController) UpdateDorama(token string, record model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}
	// TODO +adminAccessError
	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = d.repo.UpdateDorama(record)
	if err != nil {
		return fmt.Errorf("updateDorama: %w", err)
	}
	return nil
}

package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type StaffController struct {
	repo repository.IStaffRepo
	uc   controller.IUserController
}

func (s *StaffController) GetList() ([]model.Staff, error) {
	res, err := s.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("getList: %w", err)
	}
	return res, nil
}

func (s *StaffController) GetListByName(name string) ([]model.Staff, error) {
	res, err := s.repo.GetListName(name)
	if err != nil {
		return nil, fmt.Errorf("getListName: %w", err)
	}
	return res, nil
}

func (s *StaffController) GetListByDorama(idD int) ([]model.Staff, error) {
	res, err := s.repo.GetListDorama(idD)
	if err != nil {
		return nil, fmt.Errorf("getListDorama: %w", err)
	}
	return res, nil
}

func (s *StaffController) CreateStaff(token string, record model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	_, err = s.repo.CreateStaff(record)
	if err != nil {
		return fmt.Errorf("createStaff: %w", err)
	}
	return nil
}

func (s *StaffController) UpdateStaff(token string, record model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}
	err = s.repo.UpdateStaff(record)
	if err != nil {
		return fmt.Errorf("updateStaff: %w", err)
	}
	return nil
}

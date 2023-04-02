package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/interfaces"
	"errors"
)

type StaffController struct {
	repo  interfaces.IStaffRepo
	urepo interfaces.IUserRepo
}

func (s *StaffController) GetList() ([]model.Staff, error) {
	return s.repo.GetList()
}

func (s *StaffController) GetListByName(name string) ([]model.Staff, error) {
	return s.repo.GetListName(name)
}

func (s *StaffController) GetListByDorama(idD int) ([]model.Staff, error) {
	return s.repo.GetListDorama(idD)
}

func (s *StaffController) CreateStaff(username string, record model.Staff) error {
	user, err := s.urepo.GetUser(username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return errors.New("low level of access")
	}

	return s.repo.CreateStaff(record)
}

func (s *StaffController) UpdateStaff(username string, record model.Staff) error {
	user, err := s.urepo.GetUser(username)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return errors.New("low level of access")
	}
	return s.repo.UpdateStaff(record)
}

package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
)

type StaffController struct {
	repo  interfaces.IStaffRepo
	urepo interfaces.IUserRepo
}

func (s *StaffController) GetList() ([]model.Staff, error) {
	res, err := s.repo.GetList()
	return res, err
}

func (s *StaffController) GetListByName(name string) ([]model.Staff, error) {
	res, err := s.repo.GetListName(name)
	return res, err
}

func (s *StaffController) GetListByDorama(idD int) ([]model.Staff, error) {
	res, err := s.repo.GetListDorama(idD)
	return res, err
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

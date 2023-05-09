package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/sirupsen/logrus"
)

type StaffController struct {
	repo repository.IStaffRepo
	uc   controller.IUserController
	log  *logrus.Logger
}

func NewStaffController(SRepo repository.IStaffRepo, uc controller.IUserController,
	log *logrus.Logger) *StaffController {
	return &StaffController{
		repo: SRepo,
		uc:   uc,
		log:  log,
	}
}

func (s *StaffController) GetStaffList() ([]model.Staff, error) {
	res, err := s.repo.GetList()
	if err != nil {
		s.log.Warnf("get staff list err %s", err)
		return nil, fmt.Errorf("getList: %w", err)
	}
	s.log.Infof("got list staff")
	return res, nil
}

func (s *StaffController) GetListByName(name string) ([]model.Staff, error) {
	res, err := s.repo.GetListName(name)
	if err != nil {
		s.log.Warnf("get staff list by name err %s value %s", err, name)
		return nil, fmt.Errorf("getListName: %w", err)
	}
	s.log.Infof("got list staff by name value %s", name)
	return res, nil
}

func (s *StaffController) GetStaffListByDorama(idD int) ([]model.Staff, error) {
	res, err := s.repo.GetListDorama(idD)
	if err != nil {
		s.log.Warnf("get list staff by dorama err %s, value %d", err, idD)
		return nil, fmt.Errorf("getListDorama: %w", err)
	}
	s.log.Infof("get list staff by dorama value %d", idD)
	return res, nil
}

func (s *StaffController) CreateStaff(token string, record *model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		s.log.Warnf("create staff auth err %s, token %s, value %v", err, token, record)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		s.log.Warnf("create staff access err, user %s, value %v", user.Username, record)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	id, err := s.repo.CreateStaff(*record)
	if err != nil {
		s.log.Warnf("create staff err %s, user %s, value %v", err, user.Username, record)
		return fmt.Errorf("createStaff: %w", err)
	}
	record.Id = id
	s.log.Infof("create ctaff user %s, record %v", user.Username, record)
	return nil
}

func (s *StaffController) UpdateStaff(token string, record model.Staff) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		s.log.Warnf("update staff auth err %s, token %s, value %v", err, token, record)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		s.log.Warnf("update staff access err, user %s, value %v", user.Username, record)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}
	err = s.repo.UpdateStaff(record)
	if err != nil {
		s.log.Warnf("update staff err %s, user %s, value %v", err, user.Username, record)
		return fmt.Errorf("updateStaff: %w", err)
	}
	s.log.Infof("update ctaff user %s, record %v", user.Username, record)
	return nil
}

func (s *StaffController) GetStaffById(id int) (*model.Staff, error) {
	res, err := s.repo.GetStaffById(id)
	if err != nil {
		s.log.Warnf("get staff by id err %s, value %d", err, id)
		return nil, fmt.Errorf("getStaffById: %w", err)
	}
	s.log.Infof("get staff by id value %d", id)
	return res, nil
}

package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DoramaController struct {
	repo repository.IDoramaRepo
	uc   controller.IUserController
	log  *logrus.Logger
}

func NewDoramaController(DRepo repository.IDoramaRepo, uc controller.IUserController,
	log *logrus.Logger) *DoramaController {
	return &DoramaController{repo: DRepo, uc: uc, log: log}
}

func (d *DoramaController) GetAllDorama() ([]model.Dorama, error) {
	res, err := d.repo.GetList()
	if err != nil {
		d.log.Warnf("get all dorama, get list err %s", err)
		return nil, fmt.Errorf("getList: %w", err)
	}
	d.log.Infof("got all dorama")
	return res, nil
}

func (d *DoramaController) GetDoramaByName(name string) ([]model.Dorama, error) {
	res, err := d.repo.GetListName(name)
	if err != nil {
		d.log.Warnf("get dorama by name, get list name error: %s, value: %s", err, name)
		return nil, fmt.Errorf("getListName: %w", err)
	}
	d.log.Infof("got dorama by name, value %s", name)
	return res, nil
}

func (d *DoramaController) GetDoramaById(id int) (*model.Dorama, error) {
	res, err := d.repo.GetDorama(id)
	if err != nil {
		d.log.Warnf("get dorama by id error: %s, value: %d", err, id)
		return nil, fmt.Errorf("getDorama: %w", err)
	}
	d.log.Infof("got dorama by id, value %d", id)
	return res, nil
}

func (d *DoramaController) CreateDorama(token string, record *model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		d.log.Warnf("create dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		d.log.Warnf("create dorama, access err, username %s", user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	id, err := d.repo.CreateDorama(*record)
	if err != nil {
		d.log.Warnf("create dorama, err %s, username %s, value %v", err, user.Username, record)
		return fmt.Errorf("createDorama: %w", err)
	}
	record.Id = id
	d.log.Infof("created dorama, username %s, record %v", user.Username, record)
	return nil
}

func (d *DoramaController) UpdateDorama(token string, record model.Dorama) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		d.log.Warnf("update dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !user.IsAdmin {
		d.log.Warnf("update dorama, access err: %s username %s", err, user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = d.repo.UpdateDorama(record)
	if err != nil {
		d.log.Warnf("update dorama, update err: %s username %s, record %v", err, user.Username, record)
		return fmt.Errorf("updateDorama: %w", err)
	}
	d.log.Infof("updated dorama, username %s, record %v", user.Username, record)
	return nil
}

func (d *DoramaController) AddStaffToDorama(token string, idD, idS int) error {
	user, err := d.uc.AuthByToken(token)
	if err != nil {
		d.log.Warnf("update dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !user.IsAdmin {
		d.log.Warnf("update dorama, access err: %s username %s", err, user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}
	err = d.repo.AddStaff(idD, idS)
	if err != nil {
		d.log.Warnf("add staff to dorama, add err: %s values %d %d", err, idD, idS)
		return fmt.Errorf("addStaff: %w", err)
	}
	d.log.Infof("added staff to dorama, value %d %d", idD, idS)
	return nil
}

func (d *DoramaController) AddReview(token string, review model.Review) error {
	return nil
}

func (d *DoramaController) DeleteReview(token string, idD int) error {
	return nil
}

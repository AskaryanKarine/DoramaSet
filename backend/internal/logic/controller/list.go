package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ListController struct {
	repo  repository.IListRepo
	drepo repository.IDoramaRepo
	uc    controller.IUserController
	log   *logrus.Logger
}

func NewListController(LRepo repository.IListRepo, DRepo repository.IDoramaRepo,
	uc controller.IUserController, log *logrus.Logger) *ListController {
	return &ListController{
		repo:  LRepo,
		drepo: DRepo,
		uc:    uc,
		log:   log,
	}
}

func (l *ListController) CreateList(token string, record *model.List) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("create list, auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}

	record.CreatorName = user.Username

	id, err := l.repo.CreateList(*record)
	if err != nil {
		l.log.Warnf("create list err %s username %s, record %v", err, user.Username, record)
		return fmt.Errorf("createList: %w", err)
	}
	record.Id = id
	l.log.Infof("created list username %s, record %v", user.Username, record)
	return nil
}

func (l *ListController) GetUserLists(token string) ([]model.List, error) {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("get user lists err %s, token %s", err, token)
		return nil, fmt.Errorf("authToken: %w", err)
	}

	res, err := l.repo.GetUserLists(user.Username)
	if err != nil {
		l.log.Warnf("get user lists err %s, username %s", err, user.Username)
		return nil, fmt.Errorf("getUserLists: %w", err)
	}
	l.log.Infof("got user lists, username %s", user.Username)
	return res, err
}

func (l *ListController) GetPublicLists() ([]model.List, error) {

	res, err := l.repo.GetPublicLists()
	if err != nil {
		l.log.Warnf("get public list err %s", err)
		return nil, fmt.Errorf("getPublicLists: %w", err)
	}
	l.log.Infof("got public lists")
	return res, nil
}

func (l *ListController) GetListById(token string, id int) (*model.List, error) {
	res, err := l.repo.GetListId(id)
	if err != nil {
		l.log.Warnf("get list by id err %s, value %d", err, id)
		return nil, fmt.Errorf("getListById: %w", err)
	}
	if res.Type != constant.PublicList {
		user, err := l.uc.AuthByToken(token)
		if err != nil {
			l.log.Warnf("get list by id auth err %s, value %d", err, id)
			return nil, fmt.Errorf("auth: %w", err)
		}
		if user.Username != res.CreatorName {
			l.log.Warnf("get list by id access err, username %s, value %d", user.Username, id)
			return nil, fmt.Errorf("%w", errors.ErrorCreatorAccess)
		}
	}
	l.log.Infof("got list by id, value %d", id)
	return res, nil
}

func (l *ListController) AddToList(token string, idL, idD int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("add to list auth err %s, token %s, value %d %d", err, token, idL, idD)
		return fmt.Errorf("authToken: %w", err)
	}

	list, err := l.repo.GetListId(idL)
	if err != nil {
		l.log.Warnf("add to list err %s, username %s, value %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("getListId: %w", err)
	}

	if user.Username != list.CreatorName {
		l.log.Warnf("add to list access err, username %s, value %d %d", user.Username, idL, idD)
		return fmt.Errorf("%w", errors.ErrorCreatorAccess)
	}
	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		l.log.Warnf("add to list err %s, username %s, value %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("getDorama: %w", err)
	}
	err = l.repo.AddToList(idL, idD)
	if err != nil {
		l.log.Warnf("add to list err %s, username %s, value %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("addToList: %w", err)
	}
	l.log.Infof("added to list username %s, value %d %d", user.Username, idL, idD)
	return nil
}

func (l *ListController) DelFromList(token string, idL, idD int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("del from list auth err %s, token %s, value %d %d", err, token, idL, idD)
		return fmt.Errorf("authToken: %w", err)
	}

	list, err := l.GetListById(token, idL)
	if err != nil {
		l.log.Warnf("del from list err %s, username %s, values %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("getListById: %w", err)
	}

	if user.Username != list.CreatorName {
		l.log.Warnf("del from list access err, username %s, values %d %d", user.Username, idL, idD)
		return fmt.Errorf("%w", errors.ErrorCreatorAccess)
	}

	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		l.log.Warnf("del from list err %s, username %s, values %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("getDorama: %w", err)
	}
	err = l.repo.DelFromList(idL, idD)
	if err != nil {
		l.log.Warnf("del from list err %s, username %s, values %d %d", err, user.Username, idL, idD)
		return fmt.Errorf("delFromList: %w", err)
	}
	l.log.Infof("deleted from list username %s, value %d %d", user.Username, idL, idD)
	return nil
}

func (l *ListController) DelList(token string, idL int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	list, err := l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("getListId: %w", err)
	}

	if user.Username != list.CreatorName {
		return fmt.Errorf("%w", errors.ErrorCreatorAccess)
	}

	err = l.repo.DelList(idL)

	if err != nil {
		return fmt.Errorf("delList: %w", err)
	}

	return nil
}

func (l *ListController) AddToFav(token string, idL int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("add to fav auth err %s, token %s, value %d", err, token, idL)
		return fmt.Errorf("authToken: %w", err)
	}

	_, err = l.repo.GetListId(idL)
	if err != nil {
		l.log.Warnf("add to fav err %s, username %s, values %d", err, user.Username, idL)
		return fmt.Errorf("getListId: %w", err)
	}

	err = l.repo.AddToFav(idL, user.Username)
	if err != nil {
		l.log.Warnf("add to fav err %s, username %s, values %d", err, user.Username, idL)
		return fmt.Errorf("addToFav: %w", err)
	}

	l.log.Infof("added to fav username %s, value %d", user.Username, idL)
	return nil
}

func (l *ListController) GetFavList(token string) ([]model.List, error) {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		l.log.Warnf("get fav list auth err %s, token %s", err, token)
		return nil, fmt.Errorf("authToken: %w", err)
	}

	res, err := l.repo.GetFavList(user.Username)
	if err != nil {
		l.log.Warnf("get fav list err %s, username %s", err, user.Username)
		return nil, fmt.Errorf("getFavList: %w", err)
	}

	l.log.Infof("got fav list username %s", user.Username)
	return res, nil
}

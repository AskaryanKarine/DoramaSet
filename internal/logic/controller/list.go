package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type ListController struct {
	repo  repository.IListRepo
	drepo repository.IDoramaRepo
	uc    controller.IUserController
}

func (l *ListController) CreateList(token string, record model.List) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}

	record.CreatorName = user.Username

	err = l.repo.CreateList(record)
	if err != nil {
		return fmt.Errorf("createList: %w", err)
	}
	return nil
}

func (l *ListController) GetUserLists(token string) ([]model.List, error) {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return nil, err
	}

	res, err := l.repo.GetUserLists(user.Username)
	if err != nil {
		return nil, fmt.Errorf("getUserList: %w", err)
	}
	return res, err
}

func (l *ListController) GetPublicLists() ([]model.List, error) {
	res, err := l.repo.GetPublicLists()
	if err != nil {
		return nil, fmt.Errorf("getPublic: %w", err)
	}
	return res, nil
}

func (l *ListController) GetListById(id int) (*model.List, error) {
	res, err := l.repo.GetListId(id)
	if err != nil {
		return nil, fmt.Errorf("getListById: %w", err)
	}
	return res, nil
}

func (l *ListController) AddToList(token string, idL, idD int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("addToList: %w", err)
	}
	list, err := l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("addToList: %w", err)
	}
	if user.Username != list.CreatorName {
		return fmt.Errorf("addToList: no access")
	}
	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return fmt.Errorf("addToList: %w", err)
	}
	err = l.repo.AddToList(idL, idD)
	if err != nil {
		return fmt.Errorf("addToList: %w", err)
	}
	return nil
}

func (l *ListController) DelFromList(token string, idL, idD int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("delFromList: %w", err)
	}
	list, err := l.GetListById(idL)
	if err != nil {
		return fmt.Errorf("delFromList: %w", err)
	}

	if user.Username != list.CreatorName {
		return fmt.Errorf("delFromList: no access")
	}

	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return fmt.Errorf("delFromList: %w", err)
	}
	err = l.repo.DelFromList(idL, idD)
	if err != nil {
		return fmt.Errorf("delFromList: %w", err)
	}
	return nil
}

func (l *ListController) DelList(token string, idL int) error {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("delList: %w", err)
	}

	list, err := l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("delList: %w", err)
	}

	if user.Username != list.CreatorName {
		return fmt.Errorf("delList: no access")
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
		return fmt.Errorf("auth: %w", err)
	}
	_, err = l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("getListById: %w", err)
	}
	err = l.repo.AddToFav(idL, user.Username)
	if err != nil {
		return fmt.Errorf("addToFav: %w", err)
	}
	return nil
}

func (l *ListController) GetFavList(token string) ([]model.List, error) {
	user, err := l.uc.AuthByToken(token)
	if err != nil {
		return nil, fmt.Errorf("getFavList: %w", err)
	}
	res, err := l.repo.GetFavList(user.Username)
	if err != nil {
		return nil, fmt.Errorf("getFavList: %w", err)
	}
	return res, nil
}

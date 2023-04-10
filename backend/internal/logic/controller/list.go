package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
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
		return fmt.Errorf("authToken: %w", err)
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
		return nil, fmt.Errorf("authToken: %w", err)
	}

	res, err := l.repo.GetUserLists(user.Username)
	if err != nil {
		return nil, fmt.Errorf("getUserLists: %w", err)
	}
	return res, err
}

func (l *ListController) GetPublicLists() ([]model.List, error) {
	res, err := l.repo.GetPublicLists()
	if err != nil {
		return nil, fmt.Errorf("getPublicLists: %w", err)
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
		return fmt.Errorf("authToken: %w", err)
	}
	list, err := l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("getListId: %w", err)
	}
	// TODO +creatorAccessError
	if user.Username != list.CreatorName {
		return fmt.Errorf("%w", errors.ErrorCreatorAccess)
	}
	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return fmt.Errorf("getDorama: %w", err)
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
		return fmt.Errorf("authToken: %w", err)
	}
	list, err := l.GetListById(idL)
	if err != nil {
		return fmt.Errorf("getListById: %w", err)
	}
	// TODO +creatorAccessError
	if user.Username != list.CreatorName {
		return fmt.Errorf("%w", errors.ErrorCreatorAccess)
	}

	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return fmt.Errorf("getDorama: %w", err)
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
		return fmt.Errorf("authToken: %w", err)
	}

	list, err := l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("getListId: %w", err)
	}

	//TODO +creatorAccessError
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
		return fmt.Errorf("authToken: %w", err)
	}

	_, err = l.repo.GetListId(idL)
	if err != nil {
		return fmt.Errorf("getListId: %w", err)
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
		return nil, fmt.Errorf("authToken: %w", err)
	}
	res, err := l.repo.GetFavList(user.Username)
	if err != nil {
		return nil, fmt.Errorf("getFavList: %w", err)
	}
	return res, nil
}

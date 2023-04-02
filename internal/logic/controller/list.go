package controller

import (
	"DoramaSet/internal/logic/model"
	repository "DoramaSet/internal/repository/interfaces"
)

type ListController struct {
	repo  repository.IListRepo
	urepo repository.IUserRepo
	drepo repository.IDoramaRepo
}

func (l *ListController) CreateList(record model.List) error {
	_, err := l.urepo.GetUser(record.CreatorName)
	if err != nil {
		return err
	}

	return l.repo.CreateList(record)
}

func (l *ListController) GetUserLists(username string) ([]model.List, error) {
	_, err := l.urepo.GetUser(username)
	if err != nil {
		return nil, err
	}

	return l.repo.GetUserLists(username)
}

func (l *ListController) GetPublicLists() ([]model.List, error) {
	return l.repo.GetPublicLists()
}

func (l *ListController) GetListById(id int) (model.List, error) {
	return l.repo.GetListId(id)
}

func (l *ListController) AddToList(idL, idD int) error {
	_, err := l.GetListById(idL)
	if err != nil {
		return err
	}
	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return err
	}
	return l.repo.AddToList(idL, idD)
}

func (l *ListController) DelFromList(idL, idD int) error {
	_, err := l.GetListById(idL)
	if err != nil {
		return err
	}
	_, err = l.drepo.GetDorama(idD)
	if err != nil {
		return err
	}
	return l.repo.DelFromList(idL, idD)
}

func (l *ListController) DelList(idL int) error {
	_, err := l.GetListById(idL)
	if err != nil {
		return err
	}
	return l.DelList(idL)
}

func (l *ListController) AddToFav(idL int, username string) error {
	_, err := l.urepo.GetUser(username)
	if err != nil {
		return err
	}
	_, err = l.GetListById(idL)
	if err != nil {
		return err
	}
	return l.repo.AddToFav(idL, username)
}

func (l *ListController) GetFavList(username string) ([]model.List, error) {
	_, err := l.urepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	return l.repo.GetFavList(username)
}

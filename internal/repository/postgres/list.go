package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type ListRepo struct {
	db *gorm.DB
}

func (ListRepo) GetUserLists(username string) ([]model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) GetPublicLists() ([]model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) GetListId(id int) (*model.List, error) {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) CreateList(list model.List) error {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) DelList(id int) error {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) AddToList(idL, IdD int) error {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) DelFromList(idL, idD int) error {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) AddToFav(idL int, username string) error {
	//TODO implement me
	panic("implement me")
}

func (ListRepo) GetFavList(username string) ([]model.List, error) {
	//TODO implement me
	panic("implement me")
}

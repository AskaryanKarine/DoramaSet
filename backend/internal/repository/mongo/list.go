package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListRepo struct {
	db         *mongo.Database
	doramaRepo repository.IDoramaRepo
}

func NewListRepo(db *mongo.Database, DR repository.IDoramaRepo) *ListRepo {
	return &ListRepo{db, DR}
}

func (ListRepo) GetUserLists(username string) ([]model.List, error) {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) GetPublicLists() ([]model.List, error) {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) GetListId(id int) (*model.List, error) {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) CreateList(list model.List) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) DelList(id int) error {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) AddToList(idL, idD int) error {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) DelFromList(idL, idD int) error {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) AddToFav(idL int, username string) error {
	// TODO implement me
	panic("implement me")
}

func (ListRepo) GetFavList(username string) ([]model.List, error) {
	// TODO implement me
	panic("implement me")
}

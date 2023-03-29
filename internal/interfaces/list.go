package interfaces

import "DoramaSet/internal/logic/model"

type IListController interface {
	createList(record model.List) error
	getUserLists(username string) ([]model.List, error)
	getPublicLists() ([]model.List, error)
	getListById(id int) (model.List, error)
	addToList(idL, IdD int) error
	delFromList(idL, idD int) error
	delList(idL int) error
	addToFav(idL int, username string) error
	getFavList(username string) ([]model.List, error)
}

type IListRepo interface {
}

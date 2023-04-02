package interfaces

import "DoramaSet/internal/logic/model"

type IListRepo interface {
	GetUserLists(username string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	GetListId(id int) (model.List, error)
	CreateList(record model.List) error
	DelList(id int) error
	AddToList(idL, IdD int) error
	DelFromList(idL, idD int) error
	AddToFav(idL int, username string) error
	GetFavList(username string) ([]model.List, error)
}

package interfaces

import "DoramaSet/internal/logic/model"

type IListController interface {
	CreateList(record model.List) error
	GetUserLists(username string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	DetListById(id int) (model.List, error)
	AddToList(idL, idD int) error
	DelFromList(idL, idD int) error
	DelList(idL int) error
	AddToFav(idL int, username string) error
	GetFavList(username string) ([]model.List, error)
}

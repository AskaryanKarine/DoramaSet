package repository

import "DoramaSet/internal/logic/model"

type IDoramaRepo interface {
	GetList() ([]model.Dorama, error)
	GetListName(name string) ([]model.Dorama, error)
	GetDorama(id int) (*model.Dorama, error)
	CreateDorama(dorama model.Dorama) (int, error)
	UpdateDorama(dorama model.Dorama) error
	DeleteDorama(dorama model.Dorama) error
}

type IEpisodeRepo interface {
	GetList(idDorama int) ([]model.Episode, error)
	GetEpisode(id int) (*model.Episode, error)
	MarkEpisode(idEp int, username string) error
	CreateEpisode(episode model.Episode, idD int) (int, error)
	DeleteEpisode(episode model.Episode) error
}

type IListRepo interface {
	GetUserLists(username string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	GetListId(id int) (*model.List, error)
	CreateList(list model.List) error
	DelList(id int) error
	AddToList(idL, IdD int) error
	DelFromList(idL, idD int) error
	AddToFav(idL int, username string) error
	GetFavList(username string) ([]model.List, error)
}

type IPictureRepo interface {
	GetListDorama(idDorama int) ([]model.Picture, error)
	GetListStaff(idStaff int) ([]model.Picture, error)
	CreatePicture(record model.Picture, id int, tbl string) (int, error)
	DeletePicture(record model.Picture) error
}

type IStaffRepo interface {
	GetList() ([]model.Staff, error)
	GetListName(name string) ([]model.Staff, error)
	GetListDorama(idDorama int) ([]model.Staff, error)
	CreateStaff(record model.Staff) (int, error)
	UpdateStaff(record model.Staff) error
	DeleteStaff(record model.Staff) error
}

type ISubscriptionRepo interface {
	GetList() ([]model.Subscription, error)
	GetSubscription(id int) (*model.Subscription, error)
	GetSubscriptionByPrice(price int) (*model.Subscription, error)
}

type IUserRepo interface {
	GetUser(username string) (*model.User, error)
	CreateUser(record model.User) error
	UpdateUser(record model.User) error
	DeleteUser(record model.User) error
}

package controller

import "DoramaSet/internal/logic/model"

type IDoramaController interface {
	GetAll() ([]model.Dorama, error)
	GetByName(name string) ([]model.Dorama, error)
	GetById(id int) (model.Dorama, error)
	CreateDorama(token string, record model.Dorama) error
	UpdateDorama(token string, record model.Dorama) error
}

type IEpisodeController interface {
	GetEpisodeList(idD int) ([]model.Episode, error)
	GetEpisode(id int) (*model.Episode, error)
	MarkWatchingEpisode(token string, idEp int) error
}

type IListController interface {
	CreateList(token string, record model.List) error
	GetUserLists(token string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	DetListById(id int) (*model.List, error)
	AddToList(token string, idL, idD int) error
	DelFromList(token string, idL, idD int) error
	DelList(token string, idL int) error
	AddToFav(token string, idL int) error
	GetFavList(token string) ([]model.List, error)
}

type IPictureController interface {
	GetListByDorama(idD int) ([]model.Picture, error)
	GetListByStaff(idS int) ([]model.Staff, error)
	CreatePicture(token string, record model.Picture, idT int, table string) error
}

type IPointsController interface {
	EarnPointForLogin(username string) error
	EarnPoint(username string, point int) error
	PurgePoint(username string, point int) error
}

type IStaffController interface {
	GetList() ([]model.Staff, error)
	GetListByName(name string) ([]model.Staff, error)
	GetListByDorama(idD int) ([]model.Staff, error)
	CreateStaff(token string, record model.Staff) error
	UpdateStaff(token string, record model.Staff) error
}

type ISubscriptionController interface {
	GetAll() ([]model.Subscription, error)
	GetInfo(id int) (*model.Subscription, error)
	SubscribeUser(token string, id int) error
	UnsubscribeUser(token string) error
}

type IUserController interface {
	Registration(record model.User) (string, error)
	Login(username, password string) (string, error)
	UpdateActive(token string) error
	AuthByToken(token string) (*model.User, error)
}

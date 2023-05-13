package controller

import "DoramaSet/internal/logic/model"

type IDoramaController interface {
	GetAllDorama() ([]model.Dorama, error)
	GetDoramaByName(name string) ([]model.Dorama, error)
	GetDoramaById(id int) (*model.Dorama, error)
	CreateDorama(token string, record *model.Dorama) error
	UpdateDorama(token string, record model.Dorama) error
	AddStaffToDorama(token string, idD, idS int) error
}

type IEpisodeController interface {
	GetEpisodeList(idD int) ([]model.Episode, error)
	GetEpisode(id int) (*model.Episode, error)
	MarkWatchingEpisode(token string, idEp int) error
	CreateEpisode(token string, record *model.Episode, idD int) error
}

type IListController interface {
	CreateList(token string, record *model.List) error
	GetUserLists(token string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	GetListById(token string, id int) (*model.List, error)
	AddToList(token string, idL, idD int) error
	DelFromList(token string, idL, idD int) error
	DelList(token string, idL int) error
	AddToFav(token string, idL int) error
	GetFavList(token string) ([]model.List, error)
}

type IPictureController interface {
	GetListByDorama(idD int) ([]model.Picture, error)
	GetListByStaff(idS int) ([]model.Picture, error)
	CreatePicture(token string, record *model.Picture) error
	AddPictureToStaff(token string, record model.Picture, id int) error
	AddPictureToDorama(token string, record model.Picture, id int) error
}

type IPointsController interface {
	EarnPointForLogin(user *model.User) error
	EarnPoint(user *model.User, point int) error
	PurgePoint(user *model.User, point int) error
}

type IStaffController interface {
	GetStaffList() ([]model.Staff, error)
	GetListByName(name string) ([]model.Staff, error)
	GetStaffListByDorama(idD int) ([]model.Staff, error)
	GetStaffById(id int) (*model.Staff, error)
	CreateStaff(token string, record *model.Staff) error
	UpdateStaff(token string, record model.Staff) error
}

type ISubscriptionController interface {
	GetAll() ([]model.Subscription, error)
	GetInfo(id int) (*model.Subscription, error)
	SubscribeUser(token string, id int) error
	UnsubscribeUser(token string) error
	UpdateSubscribe(token string) error
}

type IUserController interface {
	Registration(record *model.User) (string, error)
	Login(username, password string) (string, error)
	UpdateActive(token string) error
	AuthByToken(token string) (*model.User, error)
	ChangeEmoji(token, emojiCode string) error
	ChangeAvatarColor(token, color string) error
}

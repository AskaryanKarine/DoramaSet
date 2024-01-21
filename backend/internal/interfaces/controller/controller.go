package controller

import (
	"DoramaSet/internal/logic/model"
	"context"
)

type IDoramaController interface {
	GetAllDorama(ctx context.Context) ([]model.Dorama, error)
	GetDoramaByName(ctx context.Context, name string) ([]model.Dorama, error)
	GetDoramaById(ctx context.Context, id int) (*model.Dorama, error)
	CreateDorama(ctx context.Context, token string, record *model.Dorama) error
	UpdateDorama(ctx context.Context, token string, record model.Dorama) error
	AddStaffToDorama(ctx context.Context, token string, idD, idS int) error
	AddReview(ctx context.Context, token string, idD int, review *model.Review) error
	DeleteReview(ctx context.Context, token string, idD int) error
}

type IEpisodeController interface {
	GetEpisodeList(ctx context.Context, idD int) ([]model.Episode, error)
	GetEpisode(ctx context.Context, id int) (*model.Episode, error)
	MarkWatchingEpisode(ctx context.Context, token string, idEp int) error
	GetWatchingEpisode(ctx context.Context, token string, idD int) ([]model.Episode, error)
	CreateEpisode(ctx context.Context, token string, record *model.Episode, idD int) error
}

type IListController interface {
	CreateList(ctx context.Context, token string, record *model.List) error
	GetUserLists(ctx context.Context, token string) ([]model.List, error)
	GetPublicLists(ctx context.Context) ([]model.List, error)
	GetListById(ctx context.Context, token string, id int) (*model.List, error)
	AddToList(ctx context.Context, token string, idL, idD int) error
	DelFromList(ctx context.Context, token string, idL, idD int) error
	DelList(ctx context.Context, token string, idL int) error
	AddToFav(ctx context.Context, token string, idL int) error
	GetFavList(ctx context.Context, token string) ([]model.List, error)
}

type IPictureController interface {
	GetListByDorama(ctx context.Context, idD int) ([]model.Picture, error)
	GetListByStaff(ctx context.Context, idS int) ([]model.Picture, error)
	CreatePicture(ctx context.Context, token string, record *model.Picture) error
	AddPictureToStaff(ctx context.Context, token string, record model.Picture, id int) error
	AddPictureToDorama(ctx context.Context, token string, record model.Picture, id int) error
}

type IPointsController interface {
	EarnPointForLogin(ctx context.Context, user *model.User) error
	EarnPoint(ctx context.Context, user *model.User, point int) error
	PurgePoint(ctx context.Context, user *model.User, point int) error
}

type IStaffController interface {
	GetStaffList(ctx context.Context) ([]model.Staff, error)
	GetListByName(ctx context.Context, name string) ([]model.Staff, error)
	GetStaffListByDorama(ctx context.Context, idD int) ([]model.Staff, error)
	GetStaffById(ctx context.Context, id int) (*model.Staff, error)
	CreateStaff(ctx context.Context, token string, record *model.Staff) error
	UpdateStaff(ctx context.Context, token string, record model.Staff) error
}

type ISubscriptionController interface {
	GetAll(ctx context.Context) ([]model.Subscription, error)
	GetInfo(ctx context.Context, id int) (*model.Subscription, error)
	SubscribeUser(ctx context.Context, token string, id int) error
	UnsubscribeUser(ctx context.Context, token string) error
	UpdateSubscribe(ctx context.Context, token string) error
}

type IUserController interface {
	Registration(ctx context.Context, record *model.User) (string, error)
	Login(ctx context.Context, username, password string) (string, error)
	UpdateActive(ctx context.Context, token string) error
	AuthByToken(ctx context.Context, token string) (*model.User, error)
	ChangeEmoji(ctx context.Context, token, emojiCode string) error
	ChangeAvatarColor(ctx context.Context, token, color string) error
	GetPublicInfo(ctx context.Context, username string) (*model.User, error)
}

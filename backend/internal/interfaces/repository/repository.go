package repository

import (
	"DoramaSet/internal/logic/model"
	"context"
)

type IDoramaRepo interface {
	GetList(ctx context.Context) ([]model.Dorama, error)
	GetListName(ctx context.Context, name string) ([]model.Dorama, error)
	GetDorama(ctx context.Context, id int) (*model.Dorama, error)
	CreateDorama(ctx context.Context, dorama model.Dorama) (int, error)
	UpdateDorama(ctx context.Context, dorama model.Dorama) error
	DeleteDorama(ctx context.Context, id int) error
	AddStaff(ctx context.Context, idD, idS int) error
	GetListByListId(ctx context.Context, idL int) ([]model.Dorama, error)
}

type IEpisodeRepo interface {
	GetList(ctx context.Context, idDorama int) ([]model.Episode, error)
	GetWatchingList(ctx context.Context, username string, idD int) ([]model.Episode, error)
	GetEpisode(ctx context.Context, id int) (*model.Episode, error)
	MarkEpisode(ctx context.Context, idEp int, username string) error
	CreateEpisode(ctx context.Context, episode model.Episode, idD int) (int, error)
	DeleteEpisode(ctx context.Context, id int) error
}

type IListRepo interface {
	GetUserLists(ctx context.Context, username string) ([]model.List, error)
	GetPublicLists(ctx context.Context) ([]model.List, error)
	GetListId(ctx context.Context, id int) (*model.List, error)
	CreateList(ctx context.Context, list model.List) (int, error)
	DelList(ctx context.Context, id int) error
	AddToList(ctx context.Context, idL, idD int) error
	DelFromList(ctx context.Context, idL, idD int) error
	AddToFav(ctx context.Context, idL int, username string) error
	GetFavList(ctx context.Context, username string) ([]model.List, error)
}

type IPictureRepo interface {
	GetListDorama(ctx context.Context, idDorama int) ([]model.Picture, error)
	GetListStaff(ctx context.Context, idStaff int) ([]model.Picture, error)
	CreatePicture(ctx context.Context, record model.Picture) (int, error)
	AddPictureToStaff(ctx context.Context, record model.Picture, id int) error
	AddPictureToDorama(ctx context.Context, record model.Picture, id int) error
	DeletePicture(ctx context.Context, id int) error
}

type IStaffRepo interface {
	GetList(ctx context.Context) ([]model.Staff, error)
	GetListName(ctx context.Context, name string) ([]model.Staff, error)
	GetListDorama(ctx context.Context, idDorama int) ([]model.Staff, error)
	CreateStaff(ctx context.Context, record model.Staff) (int, error)
	UpdateStaff(ctx context.Context, record model.Staff) error
	DeleteStaff(ctx context.Context, id int) error
	GetStaffById(ctx context.Context, id int) (*model.Staff, error)
}

type ISubscriptionRepo interface {
	GetList(ctx context.Context) ([]model.Subscription, error)
	GetSubscription(ctx context.Context, id int) (*model.Subscription, error)
	GetSubscriptionByPrice(ctx context.Context, price int) (*model.Subscription, error)
}

type IUserRepo interface {
	GetUser(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, record *model.User) error
	UpdateUser(ctx context.Context, record model.User) error
	DeleteUser(ctx context.Context, username string) error
	GetPublicInfo(ctx context.Context, username string) (*model.User, error)
}

type IReviewRepo interface {
	GetAllReview(ctx context.Context, idD int) ([]model.Review, error)
	CreateReview(ctx context.Context, idD int, record *model.Review) error
	DeleteReview(ctx context.Context, username string, idD int) error
	AggregateRate(ctx context.Context, idD int) (float64, int, error)
	GetReview(ctx context.Context, username string, idD int) (*model.Review, error)
}

type AllRepository struct {
	Dorama       IDoramaRepo
	Episode      IEpisodeRepo
	List         IListRepo
	Picture      IPictureRepo
	Staff        IStaffRepo
	Subscription ISubscriptionRepo
	User         IUserRepo
	Review       IReviewRepo
}

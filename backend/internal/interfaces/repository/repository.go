package repository

import "DoramaSet/internal/logic/model"

type IDoramaRepo interface {
	GetList() ([]model.Dorama, error)
	GetListName(name string) ([]model.Dorama, error)
	GetDorama(id int) (*model.Dorama, error)
	CreateDorama(dorama model.Dorama) (int, error)
	UpdateDorama(dorama model.Dorama) error
	DeleteDorama(id int) error
	AddStaff(idD, idS int) error
	GetListByListId(idL int) ([]model.Dorama, error)
}

type IEpisodeRepo interface {
	GetList(idDorama int) ([]model.Episode, error)
	GetWatchingList(username string, idD int) ([]model.Episode, error)
	GetEpisode(id int) (*model.Episode, error)
	MarkEpisode(idEp int, username string) error
	CreateEpisode(episode model.Episode, idD int) (int, error)
	DeleteEpisode(id int) error
}

type IListRepo interface {
	GetUserLists(username string) ([]model.List, error)
	GetPublicLists() ([]model.List, error)
	GetListId(id int) (*model.List, error)
	CreateList(list model.List) (int, error)
	DelList(id int) error
	AddToList(idL, idD int) error
	DelFromList(idL, idD int) error
	AddToFav(idL int, username string) error
	GetFavList(username string) ([]model.List, error)
}

type IPictureRepo interface {
	GetListDorama(idDorama int) ([]model.Picture, error)
	GetListStaff(idStaff int) ([]model.Picture, error)
	CreatePicture(record model.Picture) (int, error)
	AddPictureToStaff(record model.Picture, id int) error
	AddPictureToDorama(record model.Picture, id int) error
	DeletePicture(id int) error
}

type IStaffRepo interface {
	GetList() ([]model.Staff, error)
	GetListName(name string) ([]model.Staff, error)
	GetListDorama(idDorama int) ([]model.Staff, error)
	CreateStaff(record model.Staff) (int, error)
	UpdateStaff(record model.Staff) error
	DeleteStaff(id int) error
	GetStaffById(id int) (*model.Staff, error)
}

type ISubscriptionRepo interface {
	GetList() ([]model.Subscription, error)
	GetSubscription(id int) (*model.Subscription, error)
	GetSubscriptionByPrice(price int) (*model.Subscription, error)
}

type IUserRepo interface {
	GetUser(username string) (*model.User, error)
	CreateUser(record *model.User) error
	UpdateUser(record model.User) error
	DeleteUser(username string) error
}

type IReviewRepo interface {
	GetAllReview(idD int) ([]model.Review, error)
	CreateReview(idD int, record *model.Review) error
	DeleteReview(username string, idD int) error
	AggregateRate(idD int) (float64, int, error)
}

package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type DoramaRepo struct {
	db      *mongo.Database
	picRepo repository.IPictureRepo
	epRepo  repository.IEpisodeRepo
	revRepo repository.IReviewRepo
}

func NewDoramaRepo(db *mongo.Database, PR repository.IPictureRepo, ER repository.IEpisodeRepo, RR repository.IReviewRepo) *DoramaRepo {
	return &DoramaRepo{db, PR, ER, RR}
}

func (DoramaRepo) GetList() ([]model.Dorama, error) {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) GetListName(name string) ([]model.Dorama, error) {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) GetDorama(id int) (*model.Dorama, error) {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) CreateDorama(dorama model.Dorama) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) UpdateDorama(dorama model.Dorama) error {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) DeleteDorama(id int) error {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) AddStaff(idD, idS int) error {
	// TODO implement me
	panic("implement me")
}

func (DoramaRepo) GetListByListId(idL int) ([]model.Dorama, error) {
	// TODO implement me
	panic("implement me")
}

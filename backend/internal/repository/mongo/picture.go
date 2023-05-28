package mongo

import (
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type PictureRepo struct {
	db *mongo.Database
}

func NewPictureRepo(db *mongo.Database) *PictureRepo {
	return &PictureRepo{db}
}

func (PictureRepo) GetListDorama(idDorama int) ([]model.Picture, error) {
	// TODO implement me
	panic("implement me")
}

func (PictureRepo) GetListStaff(idStaff int) ([]model.Picture, error) {
	// TODO implement me
	panic("implement me")
}

func (PictureRepo) CreatePicture(record model.Picture) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (PictureRepo) AddPictureToStaff(record model.Picture, id int) error {
	// TODO implement me
	panic("implement me")
}

func (PictureRepo) AddPictureToDorama(record model.Picture, id int) error {
	// TODO implement me
	panic("implement me")
}

func (PictureRepo) DeletePicture(id int) error {
	// TODO implement me
	panic("implement me")
}

package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type PictureRepo struct {
	db *gorm.DB
}

func (PictureRepo) GetListDorama(idDorama int) ([]model.Picture, error) {
	//TODO implement me
	panic("implement me")
}

func (PictureRepo) GetListStaff(idStaff int) ([]model.Picture, error) {
	//TODO implement me
	panic("implement me")
}

func (PictureRepo) CreatePicture(record model.Picture) error {
	//TODO implement me
	panic("implement me")
}

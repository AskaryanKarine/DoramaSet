package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type DoramaRepo struct {
	db *gorm.DB
}

func (DoramaRepo) GetList() ([]model.Dorama, error) {
	//TODO implement me
	panic("implement me")
}

func (DoramaRepo) GetListName(name string) ([]model.Dorama, error) {
	//TODO implement me
	panic("implement me")
}

func (DoramaRepo) GetDorama(id int) (*model.Dorama, error) {
	//TODO implement me
	panic("implement me")
}

func (DoramaRepo) CreateDorama(dorama model.Dorama) error {
	//TODO implement me
	panic("implement me")
}

func (DoramaRepo) UpdateDorama(dorama model.Dorama) error {
	//TODO implement me
	panic("implement me")
}

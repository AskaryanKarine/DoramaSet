package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type DoramaRepo struct {
	db *gorm.DB
}

type doramaModel struct {
	ID          int
	Name        string
	Description string
	ReleaseYear int
	Status      string
	Genre       string
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

func (d DoramaRepo) CreateDorama(dorama model.Dorama) (int, error) {
	m := doramaModel{
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Status,
	}
	result := d.db.Table("dorama_set.dorama").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (DoramaRepo) UpdateDorama(dorama model.Dorama) error {
	//TODO implement me
	panic("implement me")
}

func (d DoramaRepo) DeleteDorama(dorama model.Dorama) error {
	result := d.db.Table("dorama_set.dorama").Where("id = ?", dorama.Id).Delete(&doramaModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

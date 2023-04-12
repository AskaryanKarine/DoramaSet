package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type PictureRepo struct {
	db *gorm.DB
}

func (p PictureRepo) GetListDorama(idDorama int) ([]model.Picture, error) {
	var res []model.Picture
	result := p.db.Table("dorama_set.doramapicture").Where("id_dorama = ?", idDorama).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	return res, nil
}

func (p PictureRepo) GetListStaff(idStaff int) ([]model.Picture, error) {
	var res []model.Picture
	result := p.db.Table("dorama_set.staffpicture").Where("id_staff = ?", idStaff).Find(&res)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	return res, nil
}

func (p PictureRepo) CreatePicture(record model.Picture, id int, tbl string) (int, error) {
	m := model.Picture{Description: record.Description, URL: record.URL}
	result := p.db.Table("dorama_set.picture").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}

	switch tbl {
	case "dorama":
		m1 := struct {
			IdDorama  int
			IdPicture int
		}{id, m.Id}
		result = p.db.Table("dorama_set.doramapicture").Create(&m1)
	case "staff":
		m1 := struct {
			IdStaff   int
			IdPicture int
		}{id, m.Id}
		result = p.db.Table("dorama_set.staffpicture").Create(&m1)
	}

	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}

	return m.Id, nil
}

func (p PictureRepo) DeletePicture(record model.Picture) error {
	result := p.db.Table("dorama_set.picture").Where("id = ?", record.Id).Delete(&model.Picture{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

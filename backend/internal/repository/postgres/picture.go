package postgres

import (
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type PictureRepo struct {
	db *gorm.DB
}

func NewPictureRepo(db *gorm.DB) *PictureRepo {
	return &PictureRepo{db}
}

func (p *PictureRepo) GetListDorama(idDorama int) ([]model.Picture, error) {
	var (
		res   []model.Picture
		resDB []struct {
			IdDorama  int
			IdPicture int
		}
	)

	result := p.db.Table("dorama_set.doramapicture").Where("id_dorama = ?", idDorama).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	for _, r := range resDB {
		var tmp model.Picture
		result := p.db.Table("dorama_set.picture").Where("id = ?", r.IdPicture).Take(&tmp)
		if result.Error != nil {
			return nil, fmt.Errorf("db: %w", result.Error)
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (p *PictureRepo) GetListStaff(idStaff int) ([]model.Picture, error) {
	var (
		res   []model.Picture
		resDB []struct {
			IdDorama  int
			IdPicture int
		}
	)
	result := p.db.Table("dorama_set.staffpicture").Where("id_staff = ?", idStaff).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	for _, r := range resDB {
		var tmp model.Picture
		result := p.db.Table("dorama_set.picture").Where("id = ?", r.IdPicture).Take(&tmp)
		if result.Error != nil {
			return nil, fmt.Errorf("db: %w", result.Error)
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (p *PictureRepo) CreatePicture(record model.Picture) (int, error) {
	m := model.Picture{URL: record.URL}
	result := p.db.Table("dorama_set.picture").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}

	return m.Id, nil
}

func (p *PictureRepo) AddPictureToStaff(record model.Picture, id int) error {
	m := struct {
		IdStaff   int
		IdPicture int
	}{id, record.Id}
	result := p.db.Table("dorama_set.staffpicture").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}
func (p *PictureRepo) AddPictureToDorama(record model.Picture, id int) error {
	m := struct {
		IdDorama  int
		IdPicture int
	}{id, record.Id}
	result := p.db.Table("dorama_set.doramapicture").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (p *PictureRepo) DeletePicture(id int) error {
	result := p.db.Table("dorama_set.picture").Where("id = ?", id).Delete(&model.Picture{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

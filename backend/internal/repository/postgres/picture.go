package postgres

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type PictureRepo struct {
	db *gorm.DB
}

func NewPictureRepo(db *gorm.DB) *PictureRepo {
	return &PictureRepo{db}
}

func (p *PictureRepo) GetListDorama(ctx context.Context, idDorama int) ([]model.Picture, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListDorama")
	defer span.End()
	var (
		res   []model.Picture
		resDB []struct {
			IdDorama  int
			IdPicture int
		}
	)

	result := p.db.WithContext(ctx).Table("dorama_set.doramapicture").Where("id_dorama = ?", idDorama).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	for _, r := range resDB {
		var tmp model.Picture
		result := p.db.WithContext(ctx).Table("dorama_set.picture").Where("id = ?", r.IdPicture).Take(&tmp)
		if result.Error != nil {
			return nil, fmt.Errorf("db: %w", result.Error)
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (p *PictureRepo) GetListStaff(ctx context.Context, idStaff int) ([]model.Picture, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListStaff")
	defer span.End()
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
		result := p.db.WithContext(ctx).Table("dorama_set.picture").Where("id = ?", r.IdPicture).Take(&tmp)
		if result.Error != nil {
			return nil, fmt.Errorf("db: %w", result.Error)
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (p *PictureRepo) CreatePicture(ctx context.Context, record model.Picture) (int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreatePicture")
	defer span.End()
	m := model.Picture{URL: record.URL}
	result := p.db.WithContext(ctx).Table("dorama_set.picture").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}

	return m.Id, nil
}

func (p *PictureRepo) AddPictureToStaff(ctx context.Context, record model.Picture, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo AddPictureToStaff")
	defer span.End()
	m := struct {
		IdStaff   int
		IdPicture int
	}{id, record.Id}
	result := p.db.WithContext(ctx).Table("dorama_set.staffpicture").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}
func (p *PictureRepo) AddPictureToDorama(ctx context.Context, record model.Picture, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo AddPictureToDorama")
	defer span.End()
	m := struct {
		IdDorama  int
		IdPicture int
	}{id, record.Id}
	result := p.db.WithContext(ctx).Table("dorama_set.doramapicture").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (p *PictureRepo) DeletePicture(ctx context.Context, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DeletePicture")
	defer span.End()
	result := p.db.WithContext(ctx).Table("dorama_set.picture").Where("id = ?", id).Delete(&model.Picture{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

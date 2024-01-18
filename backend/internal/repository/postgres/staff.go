package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type StaffRepo struct {
	db      *gorm.DB
	picRepo repository.IPictureRepo
}

type staffModel struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
	Type     string
}

func NewStaffRepo(db *gorm.DB, pr repository.IPictureRepo) *StaffRepo {
	return &StaffRepo{db: db, picRepo: pr}
}

func (s *StaffRepo) GetList(ctx context.Context) ([]model.Staff, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var (
		resDB []staffModel
		res   []model.Staff
	)
	result := s.db.WithContext(ctx).Table("dorama_set.staff").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s *StaffRepo) GetListName(ctx context.Context, name string) ([]model.Staff, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListName")
	defer span.End()
	var (
		resDB []staffModel
		res   []model.Staff
	)
	like := "%" + strings.TrimRight(name, "\r\n") + "%"
	result := s.db.WithContext(ctx).Table("dorama_set.staff").Where("name like ?", like).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s *StaffRepo) GetStaffById(ctx context.Context, id int) (*model.Staff, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetStaffById")
	defer span.End()
	var resDB staffModel

	result := s.db.WithContext(ctx).Table("dorama_set.staff").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	staff, err := s.picRepo.GetListStaff(ctx, resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListStaff: %w", err)
	}
	res := model.Staff{
		Id:       resDB.ID,
		Name:     resDB.Name,
		Birthday: resDB.Birthday,
		Type:     resDB.Type,
		Photo:    staff,
	}
	return &res, nil
}

func (s *StaffRepo) GetListDorama(ctx context.Context, idDorama int) ([]model.Staff, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListDorama")
	defer span.End()
	var (
		resDB []staffModel
		res   []model.Staff
	)
	result := s.db.WithContext(ctx).Table("dorama_set.staff s").Select("s.*").
		Joins("join dorama_set.doramastaff d on s.id = d.id_staff").
		Where("id_dorama = ?", idDorama).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s *StaffRepo) CreateStaff(ctx context.Context, record model.Staff) (int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreateStaff")
	defer span.End()
	m := staffModel{
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Type:     record.Type,
	}
	result := s.db.WithContext(ctx).Table("dorama_set.staff").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (s *StaffRepo) UpdateStaff(ctx context.Context, record model.Staff) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo UpdateStaff")
	defer span.End()
	m := staffModel{
		ID:       record.Id,
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Type:     record.Type,
	}
	result := s.db.WithContext(ctx).Table("dorama_set.staff").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (s *StaffRepo) DeleteStaff(ctx context.Context, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DeleteStaff")
	defer span.End()
	result := s.db.WithContext(ctx).Table("dorama_set.staff").Where("id = ?", id).Delete(&staffModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

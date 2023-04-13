package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/db_erorrs"
	"fmt"
	"gorm.io/gorm"
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
	Bio      string
	Type     string
}

func (s StaffRepo) GetList() ([]model.Staff, error) {
	var (
		resDB []staffModel
		res   []model.Staff
	)
	result := s.db.Table("dorama_set.staff").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", db_erorrs.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Bio:      r.Bio,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s StaffRepo) GetListName(name string) ([]model.Staff, error) {
	var (
		resDB []staffModel
		res   []model.Staff
	)
	result := s.db.Table("dorama_set.staff").Where("name = ?", name).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", db_erorrs.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Bio:      r.Bio,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s StaffRepo) GetStaffById(id int) (*model.Staff, error) {
	var resDB staffModel

	result := s.db.Table("dorama_set.staff").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	staff, err := s.picRepo.GetListStaff(resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListStaff: %w", err)
	}
	res := model.Staff{
		Id:       resDB.ID,
		Name:     resDB.Name,
		Birthday: resDB.Birthday,
		Type:     resDB.Type,
		Bio:      resDB.Bio,
		Photo:    staff,
	}
	return &res, nil
}

func (s StaffRepo) GetListDorama(idDorama int) ([]model.Staff, error) {
	var (
		resDB []staffModel
		res   []model.Staff
	)
	result := s.db.Raw("select s.* from dorama_set.staff s join dorama_set.doramastaff d on s.id = d.id_staff where id_dorama = ?", idDorama).Scan(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", db_erorrs.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		staff, err := s.picRepo.GetListStaff(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListStaff: %w", err)
		}
		tmp := model.Staff{
			Id:       r.ID,
			Name:     r.Name,
			Birthday: r.Birthday,
			Type:     r.Type,
			Gender:   r.Gender,
			Bio:      r.Bio,
			Photo:    staff,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (s StaffRepo) CreateStaff(record model.Staff) (int, error) {
	m := staffModel{
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Bio:      record.Bio,
		Type:     record.Type,
	}
	result := s.db.Table("dorama_set.staff").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (s StaffRepo) UpdateStaff(record model.Staff) error {
	m := staffModel{
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Bio:      record.Bio,
		Type:     record.Type,
	}
	result := s.db.Table("dorama_set.staff").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (s StaffRepo) DeleteStaff(id int) error {
	result := s.db.Table("dorama_set.staff").Where("id = ?", id).Delete(&staffModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

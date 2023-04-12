package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type StaffRepo struct {
	db *gorm.DB
}

func (StaffRepo) GetList() ([]model.Staff, error) {
	//TODO implement me
	panic("implement me")
}

func (StaffRepo) GetListName(name string) ([]model.Staff, error) {
	//TODO implement me
	panic("implement me")
}

func (StaffRepo) GetListDorama(idDorama int) ([]model.Staff, error) {
	//TODO implement me
	panic("implement me")
}

func (StaffRepo) CreateStaff(record model.Staff) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (StaffRepo) UpdateStaff(record model.Staff) error {
	//TODO implement me
	panic("implement me")
}

func (s StaffRepo) DeleteStaff(record model.Staff) error {
	panic("")
}

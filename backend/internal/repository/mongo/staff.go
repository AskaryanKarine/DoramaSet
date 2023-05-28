package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepo struct {
	db      *mongo.Database
	picRepo repository.IPictureRepo
}

func NewStaffRepo(db *mongo.Database, pr repository.IPictureRepo) *StaffRepo {
	return &StaffRepo{db: db, picRepo: pr}
}

func (StaffRepo) GetList() ([]model.Staff, error) {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) GetListName(name string) ([]model.Staff, error) {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) GetListDorama(idDorama int) ([]model.Staff, error) {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) CreateStaff(record model.Staff) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) UpdateStaff(record model.Staff) error {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) DeleteStaff(id int) error {
	// TODO implement me
	panic("implement me")
}

func (StaffRepo) GetStaffById(id int) (*model.Staff, error) {
	// TODO implement me
	panic("implement me")
}

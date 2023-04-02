package interfaces

import "DoramaSet/internal/logic/model"

type IStaffRepo interface {
	GetList() ([]model.Staff, error)
	GetListName(name string) ([]model.Staff, error)
	GetListDorama(idDorama int) ([]model.Staff, error)
	CreateStaff(record model.Staff) error
	UpdateStaff(record model.Staff) error
}

package interfaces

import "DoramaSet/internal/logic/model"

type IStaffController interface {
	GetList() ([]model.Staff, error)
	GetListByName(name string) ([]model.Staff, error)
	GetListByDorama(idD int) ([]model.Staff, error)
	CreateStaff(username string, record model.Staff) error
	UpdateStaff(username string, record model.Staff) error
}

package interfaces

import "DoramaSet/internal/logic/model"

type IPictureController interface {
	GetListByDorama(idD int) ([]model.Picture, error)
	GetListByStaff(idS int) ([]model.Staff, error)
	CreatePicture(username string, record model.Picture) error
}

type IPictureRepo interface {
	GetListDorama(idDorama int) ([]model.Picture, error)
	GetListStaff(idStaff int) ([]model.Picture, error)
	CreatePicture(record model.Picture) error
}

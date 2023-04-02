package interfaces

import "DoramaSet/internal/logic/model"

type IPictureRepo interface {
	GetListDorama(idDorama int) ([]model.Picture, error)
	GetListStaff(idStaff int) ([]model.Picture, error)
	CreatePicture(record model.Picture) error
}

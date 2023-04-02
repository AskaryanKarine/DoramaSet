package interfaces

import "DoramaSet/internal/logic/model"

type IDoramaRepo interface {
	GetList() ([]model.Dorama, error)
	GetListName(name string) ([]model.Dorama, error)
	GetDorama(id int) (model.Dorama, error)
	CreateDorama(record model.Dorama) error
	UpdateDorama(record model.Dorama) error
}

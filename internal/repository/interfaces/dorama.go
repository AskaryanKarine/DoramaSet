package interfaces

import "DoramaSet/internal/logic/model"

type IDoramaRepo interface {
	GetList() ([]model.Dorama, error)
	GetListName(name string) ([]model.Dorama, error)
	GetDorama(id int) (model.Dorama, error)
	CreateDorama(dorama model.Dorama) error
	UpdateDorama(dorama model.Dorama) error
}

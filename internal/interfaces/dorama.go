package interfaces

import "DoramaSet/internal/logic/model"

type IDoramaController interface {
	GetAll() ([]model.Dorama, error)
	GetByName(name string) ([]model.Dorama, error)
	GetById(id int) (model.Dorama, error)
	CreateDorama(username string, record model.Dorama) error
	UpdateDorama(username string, record model.Dorama) error
}

type IDoramaRepo interface {
	GetList() ([]model.Dorama, error)
	GetListName(name string) ([]model.Dorama, error)
	GetDorama(id int) (model.Dorama, error)
	CreateDorama(record model.Dorama) error
	UpdateDorama(record model.Dorama) error
}

package interfaces

import "DoramaSet/internal/logic/model"

type IDoramaController interface {
	GetAll() ([]model.Dorama, error)
	GetByName(name string) ([]model.Dorama, error)
	GetById(id int) (model.Dorama, error)
	CreateDorama(username string, record model.Dorama) error
	UpdateDorama(username string, record model.Dorama) error
}

package interfaces

import "DoramaSet/internal/logic/model"

type IUserRepo interface {
	GetUser(username string) (model.User, error)
	CreateUser(record model.User) error
	UpdateUser(record model.User) error
}

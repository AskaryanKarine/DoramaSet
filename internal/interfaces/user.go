package interfaces

import "DoramaSet/internal/logic/model"

type IUserController interface {
	Registration(record model.User) error
	Login(record model.User) error
	Logout(record model.User) error
	UpdateActive(record model.User) error
}

type IUserRepo interface {
	GetUser(username string) (model.User, error)
	CreateUser(record model.User) error
	UpdateUser(record model.User) error
}

package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
)

type UserController struct {
	repo interfaces.IUserRepo
}

func (u *UserController) Registration(record model.User) error {
	return nil
}
func (u *UserController) Login(record model.User) error {
	return nil
}

func (u *UserController) Logout(record model.User) error {
	return nil
}

func (u *UserController) UpdateActive(record model.User) error {
	return nil
}

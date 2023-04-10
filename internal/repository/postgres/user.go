package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (UserRepo) GetUser(username string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepo) CreateUser(record model.User) error {
	//TODO implement me
	panic("implement me")
}

func (UserRepo) UpdateUser(record model.User) error {
	//TODO implement me
	panic("implement me")
}

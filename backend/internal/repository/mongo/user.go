package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db       *mongo.Database
	subRepo  repository.ISubscriptionRepo
	listRepo repository.IListRepo
}

func NewUserRepo(db *mongo.Database, SR repository.ISubscriptionRepo, LR repository.IListRepo) *UserRepo {
	return &UserRepo{db, SR, LR}
}

func (UserRepo) GetUser(username string) (*model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (UserRepo) CreateUser(record *model.User) error {
	// TODO implement me
	panic("implement me")
}

func (UserRepo) UpdateUser(record model.User) error {
	// TODO implement me
	panic("implement me")
}

func (UserRepo) DeleteUser(username string) error {
	// TODO implement me
	panic("implement me")
}

func (UserRepo) GetPublicInfo(username string) (*model.User, error) {
	// TODO implement me
	panic("implement me")
}

package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo struct {
	db       *mongo.Database
	subRepo  repository.ISubscriptionRepo
	listRepo repository.IListRepo
}

type userModel struct {
	Username         string    `bson:"username"`
	SubId            int       `bson:"sub_id"`
	Password         string    `bson:"password"`
	Email            string    `bson:"email"`
	RegistrationDate time.Time `bson:"registration_date"`
	LastActive       time.Time `bson:"last_active"`
	LastSubscribe    time.Time `bson:"last_subscribe"`
	Points           int       `bson:"points"`
	IsAdmin          bool      `bson:"is_admin"`
	Emoji            string    `bson:"emoji"`
	Color            string    `bson:"color"`
}

func NewUserRepo(db *mongo.Database, SR repository.ISubscriptionRepo, LR repository.IListRepo) *UserRepo {
	return &UserRepo{db, SR, LR}
}

func (u *UserRepo) GetUser(username string) (*model.User, error) {
	var user userModel
	collection := u.db.Collection("user")
	filter := bson.D{{"username", username}}
	err := collection.FindOne(nil, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	sub, err := u.subRepo.GetSubscription(user.SubId)
	if err != nil {
		return nil, fmt.Errorf("getSubById: %w", err)
	}

	lists, err := u.listRepo.GetUserLists(user.Username)
	if err != nil {
		return nil, fmt.Errorf("getUsersLists: %w", err)
	}

	res := model.User{
		Username:      user.Username,
		Password:      user.Password,
		Email:         user.Email,
		RegData:       user.RegistrationDate,
		LastActive:    user.LastActive,
		LastSubscribe: user.LastSubscribe,
		Points:        user.Points,
		IsAdmin:       user.IsAdmin,
		Color:         user.Color,
		Emoji:         user.Emoji,
		Sub:           sub,
		Collection:    lists,
	}
	return &res, nil
}

func (u *UserRepo) CreateUser(record *model.User) error {
	collection := u.db.Collection("user")

	freeSub, err := u.subRepo.GetSubscriptionByPrice(0)
	if err != nil {
		return fmt.Errorf("getSubByPrice: %w", err)
	}
	m := userModel{
		Username:         record.Username,
		SubId:            freeSub.Id,
		Password:         record.Password,
		Email:            record.Email,
		RegistrationDate: record.RegData,
		LastActive:       record.LastActive,
		LastSubscribe:    record.LastSubscribe,
		Points:           record.Points,
		IsAdmin:          record.IsAdmin,
		Emoji:            record.Emoji,
		Color:            record.Color,
	}
	record.Sub = freeSub
	_, err = collection.InsertOne(nil, m)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}

	_, err = u.listRepo.CreateList(model.List{
		Name:        fmt.Sprintf("Просмотры %s", record.Username),
		Description: "",
		CreatorName: record.Username,
		Type:        constant.PrivateList,
	})

	if err != nil {
		return fmt.Errorf("createList: %w", err)
	}
	return nil
}

func (u *UserRepo) UpdateUser(record model.User) error {
	m := userModel{
		Username:         record.Username,
		IsAdmin:          record.IsAdmin,
		SubId:            record.Sub.Id,
		Password:         record.Password,
		Email:            record.Email,
		RegistrationDate: record.RegData,
		LastActive:       record.LastActive,
		LastSubscribe:    record.LastSubscribe,
		Points:           record.Points,
		Color:            record.Color,
		Emoji:            record.Emoji,
	}
	collection := u.db.Collection("user")
	filter := bson.D{{"username", record.Username}}
	_, err := collection.ReplaceOne(nil, filter, m)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (u *UserRepo) DeleteUser(username string) error {
	collection := u.db.Collection("user")
	filter := bson.D{{"username", username}}
	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (u *UserRepo) GetPublicInfo(username string) (*model.User, error) {
	var resDB userModel
	collection := u.db.Collection("user")
	filter := bson.D{{"username", username}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	sub, err := u.subRepo.GetSubscription(resDB.SubId)
	if err != nil {
		return nil, fmt.Errorf("getSubscription: %w", err)
	}
	m := model.User{
		Username: resDB.Username,
		Color:    resDB.Color,
		Emoji:    resDB.Emoji,
		Sub:      sub,
	}
	return &m, nil
}

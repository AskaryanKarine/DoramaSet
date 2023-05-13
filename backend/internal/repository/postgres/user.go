package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db       *gorm.DB
	subRepo  repository.ISubscriptionRepo
	listRepo repository.IListRepo
}

type userModel struct {
	Username         string `gorm:"primaryKey"`
	IsAdmin          bool
	SubId            int
	Password         string
	Email            string
	RegistrationDate time.Time
	LastActive       time.Time
	LastSubscribe    time.Time
	Points           int
	Color            string
	Emoji            string
}

func NewUserRepo(db *gorm.DB, SR repository.ISubscriptionRepo, LR repository.IListRepo) *UserRepo {
	return &UserRepo{db, SR, LR}
}

func (u *UserRepo) GetUser(username string) (*model.User, error) {
	var user *userModel
	result := u.db.Table("dorama_set.user").Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if user.Username == "" {
		return nil, nil
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
		Sub:           sub,
		Collection:    lists,
		Emoji:         user.Emoji,
	}

	return &res, nil
}

func (u *UserRepo) CreateUser(record *model.User) error {
	freeSub, err := u.subRepo.GetSubscriptionByPrice(0)
	if err != nil {
		return fmt.Errorf("getSubByPrice: %w", err)
	}

	m := userModel{
		Username:         record.Username,
		Password:         record.Password,
		Email:            record.Email,
		RegistrationDate: record.RegData,
		LastActive:       record.LastActive,
		LastSubscribe:    record.LastSubscribe,
		Points:           record.Points,
		IsAdmin:          record.IsAdmin,
		Color:            record.Color,
		Emoji:            record.Emoji,
		SubId:            freeSub.Id,
	}
	record.Sub = freeSub
	result := u.db.Table("dorama_set.user").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
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
	result := u.db.Table("dorama_set.user").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u *UserRepo) DeleteUser(username string) error {
	result := u.db.Table("dorama_set.user").Where("username = ?", username).Delete(&userModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

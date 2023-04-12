package postgres

import (
	"DoramaSet/internal/interfaces/repository"
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
	Points           int
}

func (u UserRepo) GetUser(username string) (*model.User, error) {
	var user *userModel
	//	var user []interface{}
	result := u.db.Table("dorama_set.user").Where("username = ?", username).Scan(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if user == nil {
		return nil, nil
	}

	sub, err := u.subRepo.GetSubscription(user.SubId)
	if err != nil {
		return nil, fmt.Errorf("setSubById: %w", err)
	}

	res := model.User{
		Username:   user.Username,
		Password:   user.Password,
		Email:      user.Email,
		RegData:    user.RegistrationDate,
		LastActive: user.LastActive,
		Points:     user.Points,
		IsAdmin:    user.IsAdmin,
		Sub:        sub,
		// TODO fix after list repo
		Collection: nil,
	}

	return &res, nil
}

func (u UserRepo) CreateUser(record model.User) error {
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
		Points:           record.Points,
		IsAdmin:          record.IsAdmin,
		SubId:            freeSub.Id,
	}
	// TODO create empty watching list
	result := u.db.Table("dorama_set.user").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u UserRepo) UpdateUser(record model.User) error {
	m := userModel{
		Username:         record.Username,
		IsAdmin:          record.IsAdmin,
		SubId:            record.Sub.Id,
		Password:         record.Password,
		Email:            record.Email,
		RegistrationDate: record.RegData,
		LastActive:       record.LastActive,
		Points:           record.Points,
	}
	result := u.db.Table("dorama_set.user").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u UserRepo) DeleteUser(record model.User) error {
	result := u.db.Table("dorama_set.user").Where("username = ?", record.Username).Delete(&userModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

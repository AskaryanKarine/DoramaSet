package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db      *gorm.DB
	subRepo repository.ISubscriptionRepo
}

type userModel struct {
	Username         string
	SubId            int
	Password         string
	Email            string
	RegistrationDate time.Time
	LastActive       time.Time
	Points           int
	IsAdmin          bool
}

func (u UserRepo) GetUser(username string) (*model.User, error) {
	var user *userModel
	result := u.db.Table("user").Select("user.username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
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
	result := u.db.Table("dorama_set.user").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u UserRepo) UpdateUser(record model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(record model.User) error {
	result := u.db.Table("dorama_set.user").Where("username = ?", record.Username).Delete(&userModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
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

func (u *UserRepo) GetUser(ctx context.Context, username string) (*model.User, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetUser")
	defer span.End()
	var user *userModel
	result := u.db.WithContext(ctx).Table("dorama_set.user").Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if user.Username == "" {
		return nil, nil
	}

	sub, err := u.subRepo.GetSubscription(ctx, user.SubId)
	if err != nil {
		return nil, fmt.Errorf("getSubById: %w", err)
	}

	lists, err := u.listRepo.GetUserLists(ctx, user.Username)
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

func (u *UserRepo) CreateUser(ctx context.Context, record *model.User) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreateUser")
	defer span.End()
	freeSub, err := u.subRepo.GetSubscriptionByPrice(ctx, 0)
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
	result := u.db.WithContext(ctx).Table("dorama_set.user").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}

	_, err = u.listRepo.CreateList(ctx, model.List{
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

func (u *UserRepo) UpdateUser(ctx context.Context, record model.User) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo UpdateUser")
	defer span.End()
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
	result := u.db.WithContext(ctx).Table("dorama_set.user").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, username string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DeleteUser")
	defer span.End()
	result := u.db.WithContext(ctx).Table("dorama_set.user").Where("username = ?", username).Delete(&userModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (u *UserRepo) GetPublicInfo(ctx context.Context, username string) (*model.User, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetPublicInfo")
	defer span.End()
	var resDB userModel
	result := u.db.WithContext(ctx).Table("dorama_set.user").Where("username = ?", username).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	sub, err := u.subRepo.GetSubscription(ctx, resDB.SubId)
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

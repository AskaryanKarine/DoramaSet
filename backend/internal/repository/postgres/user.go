package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db *gorm.DB
}

type userModel struct {
	Username          string
	Sub_id            int
	Password          string
	Email             string
	Registration_date time.Time
	Last_active       time.Time
	Points            int
	Is_admin          bool
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
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateUser(record model.User) error {
	//TODO implement me
	panic("implement me")
}

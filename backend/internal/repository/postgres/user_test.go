package postgres

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestUserRepo_CreateUser(t *testing.T) {
	type fields struct {
		db       *gorm.DB
		subRepo  repository.ISubscriptionRepo
		listRepo repository.IListRepo
	}
	type args struct {
		record model.User
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	sr := SubscriptionRepo{db: db}
	lr := ListRepo{db: db}
	user := model.User{Username: "qwerty"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(username string) error
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: &sr, listRepo: &lr},
			args:    args{record: user},
			wantErr: false,
			check: func(username string) error {
				res := db.Table("dorama_set.user").
					Where("username = ?", username).Take(&userModel{})
				return res.Error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:       tt.fields.db,
				subRepo:  tt.fields.subRepo,
				listRepo: tt.fields.listRepo,
			}
			if err := u.CreateUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(tt.args.record.Username); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepo_DeleteUser(t *testing.T) {
	type fields struct {
		db       *gorm.DB
		subRepo  repository.ISubscriptionRepo
		listRepo repository.IListRepo
	}
	type args struct {
		user     model.User
		username string
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	user := model.User{Username: "test"}
	sr := SubscriptionRepo{db: db}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(id string) error
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: &sr, listRepo: nil},
			args:    args{user: user, username: user.Username},
			wantErr: false,
			check: func(id string) error {
				res := db.Table("dorama_set.user").
					Where("username = ?", id).Take(&userModel{})
				return res.Error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:      tt.fields.db,
				subRepo: tt.fields.subRepo,
			}
			if err := u.DeleteUser(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(tt.args.username); (err != nil) != !tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, !tt.wantErr)
			}
		})
	}
}

func TestUserRepo_GetUser(t *testing.T) {
	type fields struct {
		db       *gorm.DB
		subRepo  repository.ISubscriptionRepo
		listRepo repository.IListRepo
	}
	type args struct {
		username string
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	sr := SubscriptionRepo{db: db}
	lr := ListRepo{db: db}
	s, _ := sr.GetSubscriptionByPrice(0)
	tm := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	user := model.User{
		Username:      "test1",
		Password:      "qwerty",
		Email:         "qwerty@gmail.com",
		RegData:       tm,
		LastActive:    tm,
		LastSubscribe: tm,
		Sub:           s,
		Collection:    nil,
		Points:        100,
		IsAdmin:       false,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: &sr, listRepo: &lr},
			args:    args{username: "test1"},
			want:    &user,
			wantErr: false,
		},
		{
			name:    "don't exists",
			fields:  fields{db: db, subRepo: &sr, listRepo: &lr},
			args:    args{username: "qerty"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:       tt.fields.db,
				subRepo:  tt.fields.subRepo,
				listRepo: tt.fields.listRepo,
			}
			got, err := u.GetUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_UpdateUser(t *testing.T) {
	type fields struct {
		db       *gorm.DB
		subRepo  repository.ISubscriptionRepo
		listRepo repository.IListRepo
	}
	type args struct {
		record model.User
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	sr := SubscriptionRepo{db: db}
	s, _ := sr.GetSubscriptionByPrice(0)
	tm := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	user := model.User{
		Username:      "test1",
		Password:      "qwerty",
		Email:         "qwerty@gmail.com",
		RegData:       tm,
		LastActive:    tm,
		LastSubscribe: tm,
		Sub:           s,
		Collection:    nil,
		Points:        120,
		IsAdmin:       false,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(username string) error
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: &sr, listRepo: nil},
			args:    args{record: user},
			wantErr: false,
			check: func(username string) error {
				var u userModel
				res := db.Table("dorama_set.user").
					Where("username = ?", username).Take(&u)
				if res.Error != nil {
					return res.Error
				}
				if u.Points != user.Points {
					return errors.New("error")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:      tt.fields.db,
				subRepo: tt.fields.subRepo,
			}
			if err := u.UpdateUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(tt.args.record.Username); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

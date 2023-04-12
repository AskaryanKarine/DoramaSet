package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
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

	sr := SubscriptionRepo{db: db}
	user := model.User{Username: "qwerty"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: sr, listRepo: nil},
			args:    args{record: user},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:      tt.fields.db,
				subRepo: tt.fields.subRepo,
			}
			if err := u.CreateUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			_ = u.DeleteUser(user)
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
		record model.User
	}
	user := model.User{Username: "qwerty"}
	sr := SubscriptionRepo{db: db}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: sr, listRepo: nil},
			args:    args{record: user},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:      tt.fields.db,
				subRepo: tt.fields.subRepo,
			}
			_ = u.CreateUser(tt.args.record)
			if err := u.DeleteUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
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
	sr := SubscriptionRepo{db: db}
	s, _ := sr.GetSubscriptionByPrice(0)
	tm := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	user := model.User{Username: "qwerty", RegData: tm, LastActive: tm, Sub: s}
	_ = UserRepo{db: db, subRepo: sr}.CreateUser(user)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: sr, listRepo: nil},
			args:    args{username: "qwerty"},
			want:    &user,
			wantErr: false,
		},
		{
			name:    "don't exists",
			fields:  fields{db: db, subRepo: sr, listRepo: nil},
			args:    args{username: "qerty"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db:      tt.fields.db,
				subRepo: tt.fields.subRepo,
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
	_ = UserRepo{db: db, subRepo: sr}.DeleteUser(user)
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
	sr := SubscriptionRepo{db: db}
	s, _ := sr.GetSubscriptionByPrice(0)
	tm := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	user := model.User{Username: "qwerty", RegData: tm, LastActive: tm, Sub: s, Points: 0}
	_ = UserRepo{db: db, subRepo: sr}.CreateUser(user)
	newUser := user
	newUser.Points += 20
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db, subRepo: sr, listRepo: nil},
			args:    args{record: newUser},
			wantErr: false,
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
		})
	}
	_ = UserRepo{db: db, subRepo: sr}.DeleteUser(user)
}

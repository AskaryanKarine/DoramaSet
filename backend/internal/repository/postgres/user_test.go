package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserRepo_CreateUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		record model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: tt.fields.db,
			}
			if err := u.CreateUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepo_GetUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: tt.fields.db,
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
		db *gorm.DB
	}
	type args struct {
		record model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepo{
				db: tt.fields.db,
			}
			if err := u.UpdateUser(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

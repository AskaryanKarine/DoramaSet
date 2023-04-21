//go:build integration

package controller

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/postgres"
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestUserController_UpdateActiveIntegrate(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())
	repo := postgres.NewSubscriptionRepo(db)
	pr := postgres.NewPictureRepo(db)
	er := postgres.NewEpisodeRepo(db)
	dr := postgres.NewDoramaRepo(db, pr, er)
	lr := postgres.NewListRepo(db, dr)
	urepo := postgres.NewUserRepo(db, repo, lr)
	pointC := PointsController{urepo}
	type fields struct {
		repo      repository.IUserRepo
		pc        controller.IPointsController
		secretKey string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(ur repository.IUserRepo) error
	}{
		{
			name: "earn point",
			fields: fields{
				repo:      urepo,
				pc:        &pointC,
				secretKey: "qwerty",
			},
			args:    args{token: getToken(model.User{Username: "test"}, "qwerty")},
			wantErr: false,
			check: func(ur repository.IUserRepo) error {
				user, err := ur.GetUser("test")
				if err != nil {
					return err
				}
				fmt.Println(user.Points, user.LastActive)
				if user.Points != 105 {
					return errors.New("error")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				repo:      tt.fields.repo,
				pc:        tt.fields.pc,
				secretKey: tt.fields.secretKey,
			}
			if err := u.UpdateActive(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActiveIntegrate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(urepo); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActiveIntegrate() error = %v, expected %v", err, tt.wantErr)
			}
		})
	}
}

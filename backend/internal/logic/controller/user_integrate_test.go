package controller

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/postgres"
	"context"
	"fmt"
	"testing"
)

func TestUserController_UpdateActiveIntegrate(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())
	repo := postgres.NewSR(db)
	pr := postgres.NewPR(db)
	er := postgres.NewER(db)
	dr := postgres.NewDR(db, pr, er)
	lr := postgres.NewLR(db, dr)
	urepo := postgres.NewUR(db, repo, lr)
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
		check   func(ur repository.IUserRepo) bool
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
			check: func(ur repository.IUserRepo) bool {
				user, err := ur.GetUser("test")
				if err != nil {
					return false
				}
				fmt.Println(user.Points, user.LastActive)
				if user.Points != 105 {
					return false
				}
				return true
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
			if f := tt.check(urepo); !f {
				t.Errorf("UpdateActiveIntegrate() error = %v, expected %v", f, true)
			}
		})
	}
}

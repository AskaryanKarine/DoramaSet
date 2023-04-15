package controller

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/postgres"
	"context"
	"testing"
)

func TestSubscriptionController_SubscribeUserIntegration(t *testing.T) {
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
	userC := UserController{repo: urepo, pc: &pointC, secretKey: "qwerty"}
	token := getToken(model.User{Username: "test"}, "qwerty")
	type fields struct {
		repo  repository.ISubscriptionRepo
		urepo repository.IUserRepo
		pc    controller.IPointsController
		uc    controller.IUserController
	}
	type args struct {
		token string
		id    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(user repository.IUserRepo) bool
	}{
		{
			name: "subscribe user",
			fields: fields{
				repo:  repo,
				urepo: urepo,
				pc:    &pointC,
				uc:    &userC,
			},
			args:    args{token: token, id: 2},
			wantErr: false,
			check: func(userRepo repository.IUserRepo) bool {
				user, err := userRepo.GetUser("test")
				if err != nil {
					return false
				}
				if user.Sub.Id != 2 {
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SubscriptionController{
				repo:  tt.fields.repo,
				urepo: tt.fields.urepo,
				pc:    tt.fields.pc,
				uc:    tt.fields.uc,
			}
			if err := s.SubscribeUser(tt.args.token, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeUserIntegration() error = %v, wantErr %v", err, tt.wantErr)
			}
			if f := tt.check(urepo); !f {
				t.Errorf("SubscribeUserIntegration() error = %v, expected %v", f, true)
			}
		})
	}
}

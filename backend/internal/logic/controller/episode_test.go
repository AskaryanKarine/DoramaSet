//go:build unit

package controller

import (
	"DoramaSet/internal/config"
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	objectMother "DoramaSet/internal/object_mother"
	"DoramaSet/internal/repository/mocks"
	"DoramaSet/internal/tracing"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
)

var resultArrayEpisode = objectMother.EpisodeMother{}.GenerateRandomEpisodeSlice(1)

func TestGetEpisodeList(t *testing.T) {
	// flag.Set("config", "../../configs/test_config.yml")
	cfg, _ := config.Init()
	_, _ = tracing.Init(cfg.OpenTelemetry.Endpoint, "test", cfg.OpenTelemetry.Ratio)
	mc := minimock.NewController(t)
	testsTable := []struct {
		name   string
		fl     EpisodeController
		arg    int
		result []model.Episode
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).GetListMock.Return(resultArrayEpisode, nil),
				uc:   nil,
			},
			arg:    1,
			result: resultArrayEpisode,
			isNeg:  false,
		},
		{
			name: "get list error",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).GetListMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		ctx := context.Background()
		ctx, span := tracing.StartSpanFromContext(ctx, fmt.Sprintf("TEST GetEpisodeList: %s", testCase.name))
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			r, err := dc.GetEpisodeList(ctx, testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetEpisodeList(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GetEpisodeList(): got: %v, expect = %v", r, testCase.result)
			}
		})
		span.End()
	}
	time.Sleep(time.Second * 6)
}

func TestGetEpisode(t *testing.T) {
	// flag.Set("config", "../../configs/test_config.yml")
	cfg, _ := config.Init()
	_, _ = tracing.Init(cfg.OpenTelemetry.Endpoint, "test", cfg.OpenTelemetry.Ratio)
	mc := minimock.NewController(t)
	testsTable := []struct {
		name   string
		fl     EpisodeController
		arg    int
		result *model.Episode
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).GetEpisodeMock.Return(&resultArrayEpisode[0], nil),
				uc:   nil,
			},
			arg:    1,
			result: &resultArrayEpisode[0],
			isNeg:  false,
		},
		{
			name: "get episode error",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).GetEpisodeMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		ctx := context.Background()
		ctx, span := tracing.StartSpanFromContext(ctx, fmt.Sprintf("TEST GetEpisode: %s", testCase.name))
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			r, err := dc.GetEpisode(ctx, testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetEpisode(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GetEpisode(): got: %v, expect = %v", r, testCase.result)
			}
		})
		span.End()
	}
	time.Sleep(time.Second * 6)
}

func TestMarkWathingEpisode(t *testing.T) {
	// flag.Set("config", "../../configs/test_config.yml")
	cfg, _ := config.Init()
	_, _ = tracing.Init(cfg.OpenTelemetry.Endpoint, "test", cfg.OpenTelemetry.Ratio)
	mc := minimock.NewController(t)
	type argument struct {
		id    int
		token string
	}
	testsTable := []struct {
		name  string
		fl    EpisodeController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).MarkEpisodeMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg: argument{
				id:    1,
				token: "",
			},
			isNeg: false,
		},
		{
			name: "mark episode error",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).MarkEpisodeMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg: argument{
				id:    1,
				token: "",
			},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: EpisodeController{
				repo: mocks.NewIEpisodeRepoMock(mc).MarkEpisodeMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				id:    1,
				token: "",
			},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		ctx := context.Background()
		ctx, span := tracing.StartSpanFromContext(ctx, fmt.Sprintf("TEST MarkWatchingEpisode: %s", testCase.name))
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			err := dc.MarkWatchingEpisode(ctx, testCase.arg.token, testCase.arg.id)
			if (err != nil) != testCase.isNeg {
				t.Errorf("MarkWatchingEpisode(): error = %v, expect = %v", err, testCase.isNeg)
			}
		})
		span.End()
	}
	time.Sleep(time.Second * 6)
}

func TestEpisodeController_CreateEpisode(t *testing.T) {
	// flag.Set("config", "../../configs/test_config.yml")
	cfg, _ := config.Init()
	_, _ = tracing.Init(cfg.OpenTelemetry.Endpoint, "test", cfg.OpenTelemetry.Ratio)
	mc := minimock.NewController(t)
	adminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(true))
	noAdminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(false))

	type fields struct {
		repo repository.IEpisodeRepo
		uc   controller.IUserController
	}
	type args struct {
		token  string
		record model.Episode
		idD    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful result",
			fields: fields{
				repo: mocks.NewIEpisodeRepoMock(mc).CreateEpisodeMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			args: args{
				token:  "",
				record: model.Episode{},
				idD:    1,
			},
			wantErr: false,
		},
		{
			name: "create error",
			fields: fields{
				repo: mocks.NewIEpisodeRepoMock(mc).CreateEpisodeMock.Return(-1, errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			args: args{
				token:  "",
				record: model.Episode{},
				idD:    1,
			},
			wantErr: true,
		},
		{
			name: "auth error",
			fields: fields{
				repo: mocks.NewIEpisodeRepoMock(mc).CreateEpisodeMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			args: args{
				token:  "",
				record: model.Episode{},
				idD:    1,
			},
			wantErr: true,
		},
		{
			name: "access error",
			fields: fields{
				repo: mocks.NewIEpisodeRepoMock(mc).CreateEpisodeMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			args: args{
				token:  "",
				record: model.Episode{},
				idD:    1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		ctx, span := tracing.StartSpanFromContext(ctx, fmt.Sprintf("TEST CreateEpisode: %s", tt.name))
		t.Run(tt.name, func(t *testing.T) {
			e := &EpisodeController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
				log:  &logrus.Logger{},
			}
			if err := e.CreateEpisode(ctx, tt.args.token, &tt.args.record, tt.args.idD); (err != nil) != tt.wantErr {
				t.Errorf("CreateEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		span.End()
	}
	time.Sleep(time.Second * 6)
}

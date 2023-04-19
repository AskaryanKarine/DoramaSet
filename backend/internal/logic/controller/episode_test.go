package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultArrayEpisode = []model.Episode{
	{
		Id:         1,
		NumSeason:  1,
		NumEpisode: 1,
	},
}

func TestGetEpisodeList(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
			}
			r, err := dc.GetEpisodeList(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetEpisodeList(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GetEpisodeList(): got: %v, expect = %v", r, testCase.result)
			}
		})
	}
}

func TestGetEpisode(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
			}
			r, err := dc.GetEpisode(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetEpisode(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GetEpisode(): got: %v, expect = %v", r, testCase.result)
			}
		})
	}
}

func TestMarkWathingEpisode(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			dc := EpisodeController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
			}
			err := dc.MarkWatchingEpisode(testCase.arg.token, testCase.arg.id)
			if (err != nil) != testCase.isNeg {
				t.Errorf("MarkWatchingEpisode(): error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestEpisodeController_CreateEpisode(t *testing.T) {
	mc := minimock.NewController(t)

	type fields struct {
		repo repository.IEpisodeRepo
		uc   controller.IUserController
	}
	type args struct {
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
				uc:   nil,
			},
			args: args{
				record: model.Episode{},
				idD:    1,
			},
			wantErr: false,
		},
		{
			name: "create error",
			fields: fields{
				repo: mocks.NewIEpisodeRepoMock(mc).CreateEpisodeMock.Return(-1, errors.New("error")),
				uc:   nil,
			},
			args: args{
				record: model.Episode{},
				idD:    1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EpisodeController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
			}
			if err := e.CreateEpisode(tt.args.record, tt.args.idD); (err != nil) != tt.wantErr {
				t.Errorf("CreateEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package controller

import (
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
			name: "negative result",
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
				t.Errorf("GetAll(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", r, testCase.result)
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
			name: "negative result",
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
				t.Errorf("GetAll(): error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(r, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", r, testCase.result)
			}
		})
	}
}

func TestMarkWathingEp(t *testing.T) {
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
			name: "negative mark",
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
			name: "negative user",
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
			err := dc.MarkWatchingEpisode(testCase.arg.id, testCase.arg.token)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetAll(): error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

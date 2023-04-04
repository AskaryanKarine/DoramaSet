package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
)

func TestEarnPointForLogin(t *testing.T) {
	mc := minimock.NewController(t)
	userFev := model.User{RegData: time.Date(2000, time.February, 29, 0, 0, 0, 0, time.UTC)}
	userReg := model.User{RegData: time.Now()}
	userReg2 := model.User{RegData: time.Now().AddDate(0, 1, 0)}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   string
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(nil),
			},
			arg:   "",
			isNeg: false,
		},
		{
			name: "negative get",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).UpdateUserMock.Return(nil),
			},
			arg:   "",
			isNeg: true,
		},
		{
			name: "negative update",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   "",
			isNeg: true,
		},
		{
			name: "successful febrary",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userFev, nil).UpdateUserMock.Return(nil),
			},
			arg:   "",
			isNeg: false,
		},
		{
			name: "successful year1",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userReg, nil).UpdateUserMock.Return(nil),
			},
			arg:   "",
			isNeg: false,
		},
		{
			name: "successful year2",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&userReg2, nil).UpdateUserMock.Return(nil),
			},
			arg:   "",
			isNeg: false,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
			}
			err := dc.EarnPointForLogin(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestPurgePoint(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		username string
		point    int
	}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "negative get",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative update",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative balance",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
			}
			err := dc.PurgePoint(testCase.arg.username, testCase.arg.point)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestEarnPoint(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		username string
		point    int
	}
	testsTable := []struct {
		name  string
		fl    PointsController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "negative get",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative update",
			fl: PointsController{
				repo: mocks.NewIUserRepoMock(mc).GetUserMock.Return(&model.User{Points: 10}, nil).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PointsController{
				repo: testCase.fl.repo,
			}
			err := dc.EarnPoint(testCase.arg.username, testCase.arg.point)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

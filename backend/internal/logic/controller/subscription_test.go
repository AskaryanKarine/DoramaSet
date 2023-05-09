package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultArraySubs = []model.Subscription{{}}

func TestGetAllSub(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     SubscriptionController
		arg    int
		result []model.Subscription
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: SubscriptionController{
				repo: mocks.NewISubscriptionRepoMock(mc).GetListMock.Return(resultArraySubs, nil),
				uc:   nil,
			},
			result: resultArraySubs,
			isNeg:  false,
		},
		{
			name: "get list error",
			fl: SubscriptionController{
				repo: mocks.NewISubscriptionRepoMock(mc).GetListMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := SubscriptionController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetAll()
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetAllDorama() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestGetInfoSub(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     SubscriptionController
		arg    int
		result *model.Subscription
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: SubscriptionController{
				repo: mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(&resultArraySubs[0], nil),
				uc:   nil,
			},
			arg:    1,
			result: &resultArraySubs[0],
			isNeg:  false,
		},
		{
			name: "get subscription error",
			fl: SubscriptionController{
				repo: mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			arg:    2,
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := SubscriptionController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetInfo(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetInfo() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetInfo() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token string
		id    int
	}
	testsTable := []struct {
		name  string
		fl    SubscriptionController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(&resultArraySubs[0], nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    mocks.NewIPointsControllerMock(mc).PurgePointMock.Return(nil),
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "update user error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(&resultArraySubs[0], nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    mocks.NewIPointsControllerMock(mc).PurgePointMock.Return(nil),
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "purge error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(&resultArraySubs[0], nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    mocks.NewIPointsControllerMock(mc).PurgePointMock.Return(errors.New("error")),
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "get subscription error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    mocks.NewIPointsControllerMock(mc).PurgePointMock.Return(nil),
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
				pc:    mocks.NewIPointsControllerMock(mc).PurgePointMock.Return(nil),
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := SubscriptionController{
				repo:  testCase.fl.repo,
				uc:    testCase.fl.uc,
				pc:    testCase.fl.pc,
				urepo: testCase.fl.urepo,
				log:   &logrus.Logger{},
			}
			err := dc.SubscribeUser(testCase.arg.token, testCase.arg.id)
			if (err != nil) != testCase.isNeg {
				t.Errorf("SubscribeUser() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token string
		id    int
	}
	testsTable := []struct {
		name  string
		fl    SubscriptionController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionByPriceMock.Return(&resultArraySubs[0], nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    nil,
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "update user error ",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionByPriceMock.Return(&resultArraySubs[0], nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    nil,
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "get subscription bu price error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionByPriceMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
				pc:    nil,
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: SubscriptionController{
				repo:  mocks.NewISubscriptionRepoMock(mc).GetSubscriptionMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
				pc:    nil,
				urepo: mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := SubscriptionController{
				repo:  testCase.fl.repo,
				uc:    testCase.fl.uc,
				pc:    testCase.fl.pc,
				urepo: testCase.fl.urepo,
				log:   &logrus.Logger{},
			}
			err := dc.UnsubscribeUser(testCase.arg.token)
			if (err != nil) != testCase.isNeg {
				t.Errorf("UnsubscribeUser() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

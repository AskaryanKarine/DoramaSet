package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultPicArray = []model.Picture{
	{
		Id:          1,
		URL:         "qwerty",
		Description: "qwerty",
	},
}

func TestGetLisByDorama(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     PictureController
		arg    int
		result []model.Picture
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).GetListDoramaMock.Return(resultPicArray, nil),
				uc:   nil,
			},
			arg:    1,
			result: resultPicArray,
			isNeg:  false,
		},
		{
			name: "negative result",
			fl: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).GetListDoramaMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			arg:    1,
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PictureController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
			}
			res, err := dc.GetListByDorama(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestGetLisByStaff(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     PictureController
		arg    int
		result []model.Picture
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).GetListStaffMock.Return(resultPicArray, nil),
				uc:   nil,
			},
			arg:    1,
			result: resultPicArray,
			isNeg:  false,
		},
		{
			name: "negative result",
			fl: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).GetListStaffMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			arg:    1,
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := PictureController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
			}
			res, err := dc.GetListByStaff(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestCreatePicture(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Picture
	}
	adminUser := model.User{IsAdmin: true}
	noadminUser := adminUser
	noadminUser.IsAdmin = false
	testToken := ""
	tests := []struct {
		name  string
		field PictureController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultPicArray[0],
			},
			isNeg: false,
		},
		{
			name: "negative auth",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultPicArray[0],
			},
			isNeg: true,
		},
		{
			name: "negative admin",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultPicArray[0],
			},
			isNeg: true,
		},
		{
			name: "update error",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultPicArray[0],
			},
			isNeg: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := PictureController{
				repo: test.field.repo,
				uc:   test.field.uc,
			}
			err := dc.CreatePicture(test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("GetByName() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

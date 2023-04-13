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

func TestGetLisByDoramaPicture(t *testing.T) {
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
			name: "get picture list by picture error ",
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
				t.Errorf("GetListByDorama() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetListByDorama() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestGetLisByStaffPicture(t *testing.T) {
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
			name: "get list picture by staff error",
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
				t.Errorf("GetListByStaff() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetListByStaff() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestCreatePicture(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token   string
		picture model.Picture
		idT     int
		table   string
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
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:   testToken,
				picture: resultPicArray[0],
				idT:     1,
				table:   "",
			},
			isNeg: false,
		},
		{
			name: "auth error",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:   testToken,
				picture: resultPicArray[0],
				idT:     1,
				table:   "",
			},
			isNeg: true,
		},
		{
			name: "admin error",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
			},
			arg: argument{
				token:   testToken,
				picture: resultPicArray[0],
				idT:     1,
				table:   "",
			},
			isNeg: true,
		},
		{
			name: "update error",
			field: PictureController{
				repo: mocks.NewIPictureRepoMock(mc).CreatePictureMock.Return(-1, errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:   testToken,
				picture: resultPicArray[0],
				idT:     1,
				table:   "",
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
			err := dc.CreatePicture(test.arg.token, test.arg.picture, 1, "")
			if (err != nil) != test.isNeg {
				t.Errorf("CreatePicture() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

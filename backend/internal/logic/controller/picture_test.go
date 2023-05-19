package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultPicArray = []model.Picture{
	{
		Id:  1,
		URL: "qwerty",
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
				log:  &logrus.Logger{},
			}
			res, err := dc.GetListByDorama(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetStaffListByDorama() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetStaffListByDorama() got: %v, expect = %v", res, testCase.result)
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
				log:  &logrus.Logger{},
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
				log:  &logrus.Logger{},
			}
			err := dc.CreatePicture(test.arg.token, &test.arg.picture)
			if (err != nil) != test.isNeg {
				t.Errorf("CreatePicture() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

func TestPictureController_AddPictureToStaff(t *testing.T) {
	mc := minimock.NewController(t)
	adminUser := model.User{IsAdmin: true}
	noadminUser := model.User{IsAdmin: false}
	type fields struct {
		repo repository.IPictureRepo
		uc   controller.IUserController
	}
	type args struct {
		token  string
		record model.Picture
		id     int
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
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: false,
		},
		{
			name: "add error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToStaffMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
		{
			name: "admin error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
		{
			name: "auth error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PictureController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
				log:  &logrus.Logger{},
			}
			if err := p.AddPictureToStaff(tt.args.token, tt.args.record, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddPictureToStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPictureController_AddPictureToDorama(t *testing.T) {
	mc := minimock.NewController(t)
	adminUser := model.User{IsAdmin: true}
	noadminUser := model.User{IsAdmin: false}
	type fields struct {
		repo repository.IPictureRepo
		uc   controller.IUserController
	}
	type args struct {
		token  string
		record model.Picture
		id     int
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
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: false,
		},
		{
			name: "add error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToDoramaMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
		{
			name: "admin error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
		{
			name: "auth error",
			fields: fields{
				repo: mocks.NewIPictureRepoMock(mc).AddPictureToDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			args:    args{"", model.Picture{}, 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PictureController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
				log:  &logrus.Logger{},
			}
			if err := p.AddPictureToDorama(tt.args.token, tt.args.record, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddPictureToDorama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

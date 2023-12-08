//go:build unit

package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	objectMother "DoramaSet/internal/object_mother"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

var resultArrayDorama = objectMother.DoramaMother{}.GenerateRandomDoramaSlice(1)

func TestGetAllDorama(t *testing.T) {
	mc := minimock.NewController(t)
	testsTable := []struct {
		name   string
		fl     DoramaController
		result []model.Dorama
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetListMock.Return(resultArrayDorama, nil),
				uc:   nil,
			},
			result: resultArrayDorama,
			isNeg:  false,
		},
		{
			name: "error get list",
			fl: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetListMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			dc := DoramaController{
				repo: test.fl.repo,
				uc:   test.fl.uc,
				log:  &logrus.Logger{},
			}
			r, err := dc.GetAllDorama()
			if (err != nil) != test.isNeg {
				t.Errorf("GetAllDorama(): error = %v, expect = %v", err, test.isNeg)
			}
			if !reflect.DeepEqual(r, test.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", r, test.result)
			}
		})
	}
}

func TestGetByNameDorama(t *testing.T) {
	mc := minimock.NewController(t)
	testTable := []struct {
		name   string
		field  DoramaController
		arg    string
		result []model.Dorama
		isNeg  bool
	}{
		{
			name: "successful result",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetListNameMock.Return(resultArrayDorama, nil),
				uc:   nil,
			},
			arg:    "qwerty",
			result: resultArrayDorama,
			isNeg:  false,
		},
		{
			name: "get by name error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetListNameMock.Return(nil, errors.New("error")),
			},
			arg:    "12345",
			result: nil,
			isNeg:  true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			dc := DoramaController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetDoramaByName(test.arg)
			if (err != nil) != test.isNeg {
				t.Errorf("GetDoramaByName() error: %v, expect: %v", err, test.isNeg)
			}
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("GetDoramaByName() got: %v, expect: %v", res, test.result)
			}
		})
	}
}

func TestByIdDorama(t *testing.T) {
	mc := minimock.NewController(t)
	tests := []struct {
		name   string
		field  DoramaController
		arg    int
		result *model.Dorama
		isNeg  bool
	}{
		{
			name: "successful result",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&resultArrayDorama[0], nil),
				uc:   nil,
			},
			arg:    1,
			result: &resultArrayDorama[0],
			isNeg:  false,
		},
		{
			name: "get by id error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
			},
			arg:    5,
			result: nil,
			isNeg:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := DoramaController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetDoramaById(test.arg)
			if (err != nil) != test.isNeg {
				t.Errorf("GetDoramaById() error: %v, expect: %v", err, test.isNeg)
			}
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("GetDoramaById() got: %v, expect: %v", res, test.result)
			}
		})
	}
}

func TestCreateDorama(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Dorama
	}
	adminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(true))
	noAdminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(false))
	testToken := ""
	tests := []struct {
		name  string
		field DoramaController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: false,
		},
		{
			name: "auth error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
		{
			name: "access error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
		{
			name: "create picture error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(-1, errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := DoramaController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			err := dc.CreateDorama(test.arg.token, &test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("CreateDorama() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

func TestUpdateDorama(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Dorama
	}
	adminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(true))
	noAdminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(false))
	testToken := ""
	tests := []struct {
		name  string
		field DoramaController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).UpdateDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: false,
		},
		{
			name: "auth error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).UpdateDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
		{
			name: "access error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).UpdateDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
		{
			name: "update error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).UpdateDoramaMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := DoramaController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			err := dc.UpdateDorama(test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("UpdateDorama() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

func TestDoramaController_AddStaffToDorama(t *testing.T) {
	mc := minimock.NewController(t)
	adminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(true))
	noAdminUser := objectMother.UserMother{}.GenerateUser(objectMother.UserWithAdmin(false))
	type fields struct {
		repo repository.IDoramaRepo
		uc   controller.IUserController
	}
	type args struct {
		token string
		idD   int
		idS   int
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
				repo: mocks.NewIDoramaRepoMock(mc).AddStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			args:    args{"", 1, 1},
			wantErr: false,
		},
		{
			name: "add error",
			fields: fields{
				repo: mocks.NewIDoramaRepoMock(mc).AddStaffMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			args:    args{"", 1, 1},
			wantErr: true,
		},
		{
			name: "auth error",
			fields: fields{
				repo: mocks.NewIDoramaRepoMock(mc).AddStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			args:    args{"", 0, 0},
			wantErr: true,
		},
		{
			name: "access error",
			fields: fields{
				repo: mocks.NewIDoramaRepoMock(mc).AddStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			args:    args{"", 0, 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoramaController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
				log:  &logrus.Logger{},
			}
			if err := d.AddStaffToDorama(tt.args.token, tt.args.idD, tt.args.idS); (err != nil) != tt.wantErr {
				t.Errorf("AddStaffToDorama() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

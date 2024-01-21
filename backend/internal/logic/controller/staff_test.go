//go:build unit

package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/object_mother"
	"DoramaSet/internal/repository/mocks"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultArrayStaff = object_mother.StaffMother{}.GenerateRandomStaffSlice(1)

func TestGetListStaff(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     StaffController
		arg    int
		result []model.Staff
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListMock.Return(resultArrayStaff, nil),
				uc:   nil,
			},
			result: resultArrayStaff,
			isNeg:  false,
		},
		{
			name: "get list staff error",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := StaffController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetStaffList(context.Background())
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetStaffList() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetStaffList() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestGetListByNameStaff(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     StaffController
		arg    string
		result []model.Staff
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListNameMock.Return(resultArrayStaff, nil),
				uc:   nil,
			},
			arg:    "",
			result: resultArrayStaff,
			isNeg:  false,
		},
		{
			name: "get list name error",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListNameMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			arg:    "",
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := StaffController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetListByName(context.Background(), testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetListByName() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetListByName() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestGetListByDoramaStaff(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     StaffController
		arg    int
		result []model.Staff
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListDoramaMock.Return(resultArrayStaff, nil),
				uc:   nil,
			},
			result: resultArrayStaff,
			isNeg:  false,
		},
		{
			name: "get list picture staff error",
			fl: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).GetListDoramaMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := StaffController{
				repo: testCase.fl.repo,
				uc:   testCase.fl.uc,
				log:  &logrus.Logger{},
			}
			res, err := dc.GetStaffListByDorama(context.Background(), testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetStaffListByDorama() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetStaffListByDorama() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestCreateStaff(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Staff
	}
	adminUser := object_mother.UserMother{}.GenerateUser(object_mother.UserWithAdmin(true))
	noAdminUser := object_mother.UserMother{}.GenerateUser(object_mother.UserWithAdmin(false))
	testToken := ""
	tests := []struct {
		name  string
		field StaffController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: false,
		},
		{
			name: "auth error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
		{
			name: "access error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(1, nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
		{
			name: "create error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(-1, errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := StaffController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			err := dc.CreateStaff(context.Background(), test.arg.token, &test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("CreateStaff() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

func TestUpdateStaff(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Staff
	}
	adminUser := object_mother.UserMother{}.GenerateUser(object_mother.UserWithAdmin(true))
	noAdminUser := object_mother.UserMother{}.GenerateUser(object_mother.UserWithAdmin(false))
	testToken := ""
	tests := []struct {
		name  string
		field StaffController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).UpdateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: false,
		},
		{
			name: "auth error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).UpdateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
		{
			name: "access error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).UpdateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(noAdminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
		{
			name: "create error",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).UpdateStaffMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := StaffController{
				repo: test.field.repo,
				uc:   test.field.uc,
				log:  &logrus.Logger{},
			}
			err := dc.UpdateStaff(context.Background(), test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("UpdateStaff() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

func TestStaffController_GetStaffById(t *testing.T) {
	mc := minimock.NewController(t)
	type fields struct {
		repo repository.IStaffRepo
		uc   controller.IUserController
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Staff
		wantErr bool
	}{
		{
			name: "successful result",
			fields: fields{
				repo: mocks.NewIStaffRepoMock(mc).GetStaffByIdMock.Return(&resultArrayStaff[0], nil),
				uc:   nil,
			},
			args:    args{1},
			want:    &resultArrayStaff[0],
			wantErr: false,
		},
		{
			name: "get by id error",
			fields: fields{
				repo: mocks.NewIStaffRepoMock(mc).GetStaffByIdMock.Return(nil, errors.New("error")),
				uc:   nil,
			},
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StaffController{
				repo: tt.fields.repo,
				uc:   tt.fields.uc,
				log:  &logrus.Logger{},
			}
			got, err := s.GetStaffById(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

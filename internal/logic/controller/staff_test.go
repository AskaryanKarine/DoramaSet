package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultArrayStaff = []model.Staff{
	{},
}

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
			name: "negative result",
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
			}
			res, err := dc.GetList()
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
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
			name: "negative result",
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
			}
			res, err := dc.GetListByName(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
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
			name: "negative result",
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

func TestCreateStaff(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		dorama model.Staff
	}
	adminUser := model.User{IsAdmin: true}
	noadminUser := adminUser
	noadminUser.IsAdmin = false
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
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: false,
		},
		{
			name: "negative auth",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: true,
		},
		{
			name: "negative admin",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
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
				repo: mocks.NewIStaffRepoMock(mc).CreateStaffMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
			}
			err := dc.CreateStaff(test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("GetByName() error: %v, expect: %v", err, test.isNeg)
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
	adminUser := model.User{IsAdmin: true}
	noadminUser := adminUser
	noadminUser.IsAdmin = false
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
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayStaff[0],
			},
			isNeg: false,
		},
		{
			name: "negative auth",
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
			name: "negative admin",
			field: StaffController{
				repo: mocks.NewIStaffRepoMock(mc).UpdateStaffMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
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
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
			}
			err := dc.UpdateStaff(test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("GetByName() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}

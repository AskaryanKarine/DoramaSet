package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
)

var resultArrayDorama = []model.Dorama{
	{
		Id:          1,
		Name:        "qwerty",
		Description: "qwerty",
		Genre:       "qwerty",
		Status:      "qwerty",
		ReleaseYear: 2000,
		Posters:     nil,
		Episodes:    nil,
	},
}

// dorama controller
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
				repo: mocks.NewIDoramaRepoMock(mc).GetListMock.Return(resultArrayDorama, nil), //GetDoramaMock.Return(res, nil),
				uc:   nil,
			},
			result: resultArrayDorama,
			isNeg:  false,
		},
		{
			name: "error get list",
			fl: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).GetListMock.Return(nil, errors.New("error")), //GetDoramaMock.Return(res, nil),
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
			}
			r, err := dc.GetAll()
			if (err != nil) != test.isNeg {
				t.Errorf("GetAll(): error = %v, expect = %v", err, test.isNeg)
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
			}
			res, err := dc.GetByName(test.arg)
			if (err != nil) != test.isNeg {
				t.Errorf("GetByName() error: %v, expect: %v", err, test.isNeg)
			}
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("GetByName() got: %v, expect: %v", res, test.result)
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
			}
			res, err := dc.GetById(test.arg)
			if (err != nil) != test.isNeg {
				t.Errorf("GetById() error: %v, expect: %v", err, test.isNeg)
			}
			if !reflect.DeepEqual(res, test.result) {
				t.Errorf("GetById() got: %v, expect: %v", res, test.result)
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
	adminUser := model.User{
		Username:   "qwerty",
		Password:   "qwerty",
		Email:      "qwerty",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    true,
		Sub:        nil,
		Collection: nil,
	}
	noadminUser := adminUser
	noadminUser.IsAdmin = false
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
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(nil),
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
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(nil),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
			},
			arg: argument{
				token:  testToken,
				dorama: resultArrayDorama[0],
			},
			isNeg: true,
		},
		{
			name: "create dorama error",
			field: DoramaController{
				repo: mocks.NewIDoramaRepoMock(mc).CreateDoramaMock.Return(errors.New("error")),
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
			}
			err := dc.CreateDorama(test.arg.token, test.arg.dorama)
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
	adminUser := model.User{
		Username:   "qwerty",
		Password:   "qwerty",
		Email:      "qwerty",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    true,
		Sub:        nil,
		Collection: nil,
	}
	noadminUser := adminUser
	noadminUser.IsAdmin = false
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
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&noadminUser, nil),
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
				uc:   mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&adminUser, nil),
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
			}
			err := dc.UpdateDorama(test.arg.token, test.arg.dorama)
			if (err != nil) != test.isNeg {
				t.Errorf("UpdateDorama() error: %v, expect: %v", err, test.isNeg)
			}
		})
	}
}
package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/gojuno/minimock/v3"
)

var resultArrayList = []model.List{
	{
		Id:          1,
		Name:        "qwerty",
		Description: "qwerty",
		CreatorName: "qwerty",
		Type:        "qwerty",
		Doramas:     nil,
	},
}

func TestCreateList(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token  string
		record model.List
	}
	testsTable := []struct {
		name  string
		fl    ListController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg: argument{
				token:  "",
				record: model.List{},
			},
			isNeg: false,
		},
		{
			name: "negative create",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg: argument{
				token:  "",
				record: model.List{},
			},
			isNeg: true,
		},
		{
			name: "negative user",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg: argument{
				token:  "",
				record: model.List{},
			},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			err := dc.CreateList(testCase.arg.token, testCase.arg.record)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestGetUserList(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     ListController
		arg    string
		result []model.List
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetUserListsMock.Return(resultArrayList, nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:    "",
			isNeg:  false,
			result: resultArrayList,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetUserListsMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			result: nil,
			isNeg:  true,
		},
		{
			name: "negative get lists",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetUserListsMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			res, err := dc.GetUserLists(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetByName() got: %v, expect: %v", res, testCase.result)
			}
		})
	}
}

func TestGetPublicLists(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     ListController
		result []model.List
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetPublicListsMock.Return(resultArrayList, nil),
				drepo: nil,
				uc:    nil,
			},
			isNeg:  false,
			result: resultArrayList,
		},
		{
			name: "negative result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetPublicListsMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    nil,
			},
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			res, err := dc.GetPublicLists()
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetByName() got: %v, expect: %v", res, testCase.result)
			}
		})
	}
}

func TestGetListById(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     ListController
		arg    int
		result *model.List
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&resultArrayList[0], nil),
				drepo: nil,
				uc:    nil,
			},
			arg:    1,
			isNeg:  false,
			result: &resultArrayList[0],
		},
		{
			name: "negative result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    nil,
			},
			arg:    1,
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			res, err := dc.GetListById(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetByName() got: %v, expect: %v", res, testCase.result)
			}
		})
	}
}

func TestGetFavList(t *testing.T) {
	mc := minimock.NewController(t)

	testsTable := []struct {
		name   string
		fl     ListController
		arg    string
		result []model.List
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetFavListMock.Return(resultArrayList, nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:    "",
			result: resultArrayList,
			isNeg:  false,
		},
		{
			name: "negative get",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetFavListMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:    "",
			result: nil,
			isNeg:  true,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetFavListMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:    "",
			result: nil,
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			res, err := dc.GetFavList(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetByName() got: %v, expect: %v", res, testCase.result)
			}
		})
	}
}

func TestAddToList(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token    string
		id1, id2 int
	}
	testsTable := []struct {
		name  string
		fl    ListController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: false,
		},
		{
			name: "negative add",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(errors.New("error")),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative get dorama",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative get dorama",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative access",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{CreatorName: "ertyu"}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			err := dc.AddToList(testCase.arg.token, testCase.arg.id1, testCase.arg.id2)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestDelFromList(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token    string
		id1, id2 int
	}
	testsTable := []struct {
		name  string
		fl    ListController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: false,
		},
		{
			name: "negative add",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(errors.New("error")),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative get dorama",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative get dorama",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "negative access",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{CreatorName: "ertyu"}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			err := dc.DelFromList(testCase.arg.token, testCase.arg.id1, testCase.arg.id2)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestDelList(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token string
		id1   int
	}
	testsTable := []struct {
		name  string
		fl    ListController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "negative del",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelListMock.Return(errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative get list",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative access",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{CreatorName: "ertyu"}, nil).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			err := dc.DelList(testCase.arg.token, testCase.arg.id1)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestAddToFav(t *testing.T) {
	mc := minimock.NewController(t)
	type argument struct {
		token string
		id1   int
	}
	testsTable := []struct {
		name  string
		fl    ListController
		arg   argument
		isNeg bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToFavMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: false,
		},
		{
			name: "negative add",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToFavMock.Return(errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative auth",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToFavMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "negative get list",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).AddToFavMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := ListController{
				repo:  testCase.fl.repo,
				drepo: testCase.fl.drepo,
				uc:    testCase.fl.uc,
			}
			err := dc.AddToFav(testCase.arg.token, testCase.arg.id1)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

package controller

import (
	"DoramaSet/internal/logic/constant"
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
		Type:        constant.PublicType,
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
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(1, nil),
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
			name: "create list error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(-1, errors.New("error")),
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
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).CreateListMock.Return(1, nil),
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
				t.Errorf("CreateList() error = %v, expect = %v", err, testCase.isNeg)
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
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetUserListsMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			result: nil,
			isNeg:  true,
		},
		{
			name: "get user lists error",
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
				t.Errorf("GetUserLists() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetUserLists() got: %v, expect: %v", res, testCase.result)
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
			name: "get public list error",
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
				t.Errorf("GetPublicLists() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetPublicLists() got: %v, expect: %v", res, testCase.result)
			}
		})
	}
}

func TestGetListById(t *testing.T) {
	mc := minimock.NewController(t)

	type argument struct {
		token string
		id    int
	}
	testsTable := []struct {
		name   string
		fl     ListController
		arg    argument
		result *model.List
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&resultArrayList[0], nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, nil),
			},
			arg:    argument{"", 1},
			isNeg:  false,
			result: &resultArrayList[0],
		},
		{
			name: "get list id error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")),
				drepo: nil,
				uc:    nil,
			},
			arg:    argument{"", 1},
			result: nil,
			isNeg:  true,
		},
		{
			name: "get private list",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{Type: "private", CreatorName: "qwerty"}, nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, nil),
			},
			arg:    argument{"", 1},
			result: &model.List{Type: "private", CreatorName: "qwerty"},
			isNeg:  false,
		},
		{
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{Type: "private", CreatorName: "qwerty"}, nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwerty"}, errors.New("error")),
			},
			arg:    argument{"", 1},
			result: nil,
			isNeg:  true,
		},
		{
			name: "access error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{Type: "private", CreatorName: "qwerty"}, nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{Username: "qwe"}, nil),
			},
			arg:    argument{"", 1},
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
			res, err := dc.GetListById(testCase.arg.token, testCase.arg.id)
			if (err != nil) != testCase.isNeg {
				t.Errorf("GetListById() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetListById() got: %v, expect: %v", res, testCase.result)
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
			name: "get favorite list error",
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
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetFavListMock.Return(resultArrayList, nil),
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
				t.Errorf("GetFavList() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GetFavList() got: %v, expect: %v", res, testCase.result)
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
			name: "add to list error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(errors.New("error")),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "get picture error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "get list id error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).AddToListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "access error",
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
				t.Errorf("AddToList() error = %v, expect = %v", err, testCase.isNeg)
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
			name: "add from list error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(errors.New("error")),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "get picture error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, errors.New("error")),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(nil, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "get list id error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).DelFromListMock.Return(nil),
				drepo: mocks.NewIDoramaRepoMock(mc).GetDoramaMock.Return(&model.Dorama{}, nil),
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1, 1},
			isNeg: true,
		},
		{
			name: "access error",
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
				t.Errorf("DelFromList() error = %v, expect = %v", err, testCase.isNeg)
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
			name: "delete list error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelListMock.Return(errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "get list id error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(nil, errors.New("error")).DelListMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "access error",
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
				t.Errorf("DelList() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestAddToFavList(t *testing.T) {
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
			name: "add to fav error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToFavMock.Return(errors.New("error")),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(&model.User{}, nil),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "auth error",
			fl: ListController{
				repo:  mocks.NewIListRepoMock(mc).GetListIdMock.Return(&model.List{}, nil).AddToFavMock.Return(nil),
				drepo: nil,
				uc:    mocks.NewIUserControllerMock(mc).AuthByTokenMock.Return(nil, errors.New("error")),
			},
			arg:   argument{"", 1},
			isNeg: true,
		},
		{
			name: "get list id error",
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
				t.Errorf("AddToFav() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

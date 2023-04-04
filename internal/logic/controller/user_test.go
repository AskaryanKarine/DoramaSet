package controller

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/repository/mocks"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func getToken(newUser model.User, secretKey string) string {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        newUser.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(secretKey))

	return ss
}

func TestRegistration(t *testing.T) {
	mc := minimock.NewController(t)
	correctUser := model.User{
		Username:   "123456789",
		Password:   "123456789",
		Email:      "mail@gmail.com",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    false,
		Sub:        nil,
		Collection: nil,
	}
	shortlogin := correctUser
	shortpasword := correctUser
	shortpasword.Password = "qw"
	shortlogin.Username = "qw"
	wrongemail := correctUser
	wrongemail.Email = "qw"
	secretKey := "qwerty"
	testsTable := []struct {
		name   string
		fl     UserController
		arg    model.User
		result string
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    correctUser,
			result: getToken(correctUser, secretKey),
			isNeg:  false,
		},
		{
			name: "exists user",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "error get user",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "short login",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    shortlogin,
			result: "",
			isNeg:  true,
		},
		{
			name: "short password",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    shortpasword,
			result: "",
			isNeg:  true,
		},
		{
			name: "wrong email",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    wrongemail,
			result: "",
			isNeg:  true,
		},
		{
			name: "error create",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(errors.New("error")),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "error earn",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(errors.New("error")),
				secretKey: secretKey,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:      testCase.fl.repo,
				pc:        testCase.fl.pc,
				secretKey: testCase.fl.secretKey,
			}
			res, err := dc.Registration(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	mc := minimock.NewController(t)
	myString := "1"
	hash, _ := bcrypt.GenerateFromPassword([]byte(myString), bcrypt.DefaultCost)

	correctUser := model.User{
		Username:   "123456789",
		Password:   string(hash),
		Email:      "mail@gmail.com",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    false,
		Sub:        nil,
		Collection: nil,
	}
	type argument struct {
		login, password string
	}
	secretKey := "qwerty"
	testsTable := []struct {
		name   string
		fl     UserController
		arg    argument
		result string
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).UpdateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    argument{"123456789", myString},
			result: getToken(correctUser, secretKey),
			isNeg:  false,
		},
		{
			name: "wrong password",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).CreateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    argument{"123456789", "123456"},
			result: "",
			isNeg:  true,
		},
		{
			name: "error get",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).UpdateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    argument{"123456789", "123456"},
			result: "",
			isNeg:  true,
		},
		{
			name: "error earn",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).UpdateUserMock.Return(nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(errors.New("error")),
				secretKey: secretKey,
			},
			arg:    argument{"123456789", "1"},
			result: "",
			isNeg:  true,
		},
		{
			name: "error update",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).UpdateUserMock.Return(errors.New("error")),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    argument{"123456789", "1"},
			result: "",
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:      testCase.fl.repo,
				pc:        testCase.fl.pc,
				secretKey: testCase.fl.secretKey,
			}
			res, err := dc.Login(testCase.arg.login, testCase.arg.password)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestUpdateActive(t *testing.T) {
	mc := minimock.NewController(t)

	correctUser := model.User{
		Username:   "123456789",
		Password:   "123456789",
		Email:      "mail@gmail.com",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    false,
		Sub:        nil,
		Collection: nil,
	}
	activeUser := correctUser
	activeUser.LastActive = activeUser.LastActive.AddDate(0, 0, -1)
	secretKey := "qwerty"
	testsTable := []struct {
		name  string
		fl    UserController
		arg   string
		isNeg bool
	}{
		{
			name: "successful result",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(&correctUser, nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:   getToken(correctUser, secretKey),
			isNeg: false,
		},
		{
			name: "negative auth",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(nil, errors.New("error")),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:   getToken(correctUser, secretKey),
			isNeg: true,
		},
		{
			name: "update points",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(&activeUser, nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:   getToken(correctUser, secretKey),
			isNeg: false,
		},
		{
			name: "error update",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(&activeUser, nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:   getToken(correctUser, secretKey),
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:      testCase.fl.repo,
				pc:        testCase.fl.pc,
				secretKey: testCase.fl.secretKey,
			}
			err := dc.UpdateActive(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestAuthByToken(t *testing.T) {
	mc := minimock.NewController(t)
	correctUser := model.User{
		Username:   "123456789",
		Password:   "123456789",
		Email:      "mail@gmail.com",
		RegData:    time.Now(),
		LastActive: time.Now(),
		Points:     0,
		IsAdmin:    false,
		Sub:        nil,
		Collection: nil,
	}
	secretKey := "qwerty"
	testsTable := []struct {
		name   string
		fl     UserController
		arg    string
		result *model.User
		isNeg  bool
	}{
		{
			name: "successful result",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    getToken(correctUser, secretKey),
			isNeg:  false,
			result: &correctUser,
		},
		{
			name: "error update",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(nil, errors.New("error")),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:    getToken(correctUser, secretKey),
			isNeg:  true,
			result: nil,
		},
		{
			name: "invalid token",
			fl: UserController{
				repo:      mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(&correctUser, nil),
				pc:        mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey: secretKey,
			},
			arg:   getToken(correctUser, "12345"),
			isNeg: true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:      testCase.fl.repo,
				pc:        testCase.fl.pc,
				secretKey: testCase.fl.secretKey,
			}
			res, err := dc.AuthByToken(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("GotAll(): got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

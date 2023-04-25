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

func getToken(newUser model.User, secretKey string, tokenExp time.Duration) string {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        newUser.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(secretKey))

	return ss
}

func TestRegistrationUser(t *testing.T) {
	tokenExpiration := time.Hour * 700
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
	shortLogin := correctUser
	shortPassword := correctUser
	shortPassword.Password = "qw"
	shortLogin.Username = "qw"
	wrongEmail := correctUser
	wrongEmail.Email = "qw"
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
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil).UpdateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: correctUser.Username,
			isNeg:  false,
		},
		{
			name: "exists user error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "get user error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "short login error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    shortLogin,
			result: "",
			isNeg:  true,
		},
		{
			name: "short password error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    shortPassword,
			result: "",
			isNeg:  true,
		},
		{
			name: "wrong email error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    wrongEmail,
			result: "",
			isNeg:  true,
		},
		{
			name: "create error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(errors.New("error")),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "earn error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(errors.New("error")),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
		{
			name: "update error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, nil).CreateUserMock.Return(nil).UpdateUserMock.Return(errors.New("error")),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    correctUser,
			result: "",
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:            testCase.fl.repo,
				pc:              testCase.fl.pc,
				secretKey:       testCase.fl.secretKey,
				tokenExpiration: testCase.fl.tokenExpiration,
				loginLen:        testCase.fl.loginLen,
				passwordLen:     testCase.fl.passwordLen,
			}
			res, err := dc.Registration(testCase.arg)
			if (err != nil) != testCase.isNeg {
				t.Errorf("Registration() error = %v, expect = %v", err, testCase.isNeg)
			}
			var claims jwt.RegisteredClaims
			_, err = jwt.ParseWithClaims(res, &claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if (err != nil) != testCase.isNeg && claims.ID != testCase.result {
				t.Errorf("Registration(): got: %v, expect = %v", claims.ID, testCase.result)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	tokenExpiration := time.Hour * 700
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
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).UpdateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    argument{"123456789", myString},
			result: correctUser.Username,
			isNeg:  false,
		},
		{
			name: "wrong password error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).CreateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    argument{"123456789", "123456"},
			result: "",
			isNeg:  true,
		},
		{
			name: "get user error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(nil, errors.New("error")).UpdateUserMock.Return(nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    argument{"123456789", "123456"},
			result: "",
			isNeg:  true,
		},
		{
			name: "update error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil).UpdateUserMock.Return(errors.New("error")),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    argument{"123456789", "1"},
			result: "",
			isNeg:  true,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			dc := UserController{
				repo:            testCase.fl.repo,
				pc:              testCase.fl.pc,
				secretKey:       testCase.fl.secretKey,
				tokenExpiration: testCase.fl.tokenExpiration,
				loginLen:        testCase.fl.loginLen,
				passwordLen:     testCase.fl.passwordLen,
			}
			res, err := dc.Login(testCase.arg.login, testCase.arg.password)
			if (err != nil) != testCase.isNeg {
				t.Errorf("Login() error = %v, expect = %v", err, testCase.isNeg)
			}
			var claims jwt.RegisteredClaims
			_, err = jwt.ParseWithClaims(res, &claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if (err != nil) != testCase.isNeg || !reflect.DeepEqual(claims.ID, testCase.result) {
				t.Errorf("Login() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

func TestUpdateActive(t *testing.T) {
	mc := minimock.NewController(t)
	tokenExpiration := time.Hour * 700
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
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(&correctUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, secretKey, tokenExpiration),
			isNeg: false,
		},
		{
			name: "auth error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(nil, errors.New("error")),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, secretKey, tokenExpiration),
			isNeg: true,
		},
		{
			name: "earn point for login error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(&activeUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(errors.New("error")),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, secretKey, tokenExpiration),
			isNeg: true,
		},
		{
			name: "update points",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(nil).GetUserMock.Return(&activeUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, secretKey, tokenExpiration),
			isNeg: false,
		},
		{
			name: "update user error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(&activeUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, secretKey, tokenExpiration),
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
				t.Errorf("UpdateActive() error = %v, expect = %v", err, testCase.isNeg)
			}
		})
	}
}

func TestAuthByToken(t *testing.T) {
	tokenExpiration := time.Hour * 700
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
				repo:            mocks.NewIUserRepoMock(mc).GetUserMock.Return(&correctUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    getToken(correctUser, secretKey, tokenExpiration),
			isNeg:  false,
			result: &correctUser,
		},
		{
			name: "update error",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(nil, errors.New("error")),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:    getToken(correctUser, secretKey, tokenExpiration),
			isNeg:  true,
			result: nil,
		},
		{
			name: "invalid token",
			fl: UserController{
				repo:            mocks.NewIUserRepoMock(mc).UpdateUserMock.Return(errors.New("error")).GetUserMock.Return(&correctUser, nil),
				pc:              mocks.NewIPointsControllerMock(mc).EarnPointForLoginMock.Return(nil),
				secretKey:       secretKey,
				tokenExpiration: tokenExpiration,
				loginLen:        5,
				passwordLen:     8,
			},
			arg:   getToken(correctUser, "12345", tokenExpiration),
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
				t.Errorf("AuthByToken() error = %v, expect = %v", err, testCase.isNeg)
			}
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("AuthByToken() got: %v, expect = %v", res, testCase.result)
			}
		})
	}
}

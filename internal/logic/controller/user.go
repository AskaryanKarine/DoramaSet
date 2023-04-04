package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo      interfaces.IUserRepo
	pc        interfaces.IPointsController
	secretKey string
}

const (
	loginLen        = 5
	passwordLen     = 8
	tokenExpiration = time.Hour * 700
)

func (u *UserController) Registration(newUser model.User) (string, error) {
	res, err := u.repo.GetUser(newUser.Username)
	if err != nil {
		return "", fmt.Errorf("registration: %w", err)
	}

	if res != nil {
		return "", errors.New("registration: user already exists")
	}

	if len(newUser.Username) < loginLen {
		return "", fmt.Errorf("registration: login must be more then %d symbols", loginLen)
	}

	if len(newUser.Password) < passwordLen {
		return "", fmt.Errorf("registration: password must be more then %d symbols", loginLen)
	}

	_, err = mail.ParseAddress(newUser.Email)
	if err != nil {
		return "", fmt.Errorf("registration: invalid email")
	}

	newUser.RegData = time.Now()
	newUser.LastActive = time.Now().Add(-time.Hour * 24)

	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	newUser.Password = string(hash)

	err = u.repo.CreateUser(newUser)
	if err != nil {
		return "", fmt.Errorf("registration: %w", err)
	}

	err = u.pc.EarnPointForLogin(newUser.Username)
	if err != nil {
		return "", fmt.Errorf("earnPoint: %w", err)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        newUser.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(u.secretKey))

	return ss, nil
}

func (u *UserController) Login(username, password string) (string, error) {
	user, err := u.repo.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("login: wrong login or password")
	}

	err = u.pc.EarnPointForLogin(username)
	if err != nil {
		return "", fmt.Errorf("earnPoint: %w", err)
	}

	user.LastActive = time.Now()
	err = u.repo.UpdateUser(*user)
	if err != nil {
		return "", fmt.Errorf("update: %w", err)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(u.secretKey))

	return ss, nil
}

func eqDate(date1, date2 time.Time) bool {
	d1, m1, y1 := date1.Date()
	d2, m2, y2 := date2.Date()

	if d1 != d2 || m1 != m2 || y1 != y2 {
		return false
	}
	return true
}

func (u *UserController) UpdateActive(token string) error {
	user, err := u.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("updateActive: %w", err)
	}
	if !eqDate(user.LastActive, time.Now()) {
		u.pc.EarnPointForLogin(user.Username)
	}

	user.LastActive = time.Now()
	err = u.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateActive: %w", err)
	}
	return nil
}

func (u *UserController) AuthByToken(token string) (*model.User, error) {
	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("AuthToken: %w", err)
	}

	user, err := u.repo.GetUser(claims.ID)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return user, nil
}

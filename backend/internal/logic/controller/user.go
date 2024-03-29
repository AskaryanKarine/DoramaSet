package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo            repository.IUserRepo
	pc              controller.IPointsController
	secretKey       string
	loginLen        int
	passwordLen     int
	tokenExpiration time.Duration
	log             *logrus.Logger
}

func NewUserController(UR repository.IUserRepo, pc controller.IPointsController,
	secretKey string, loginLen, passwordLen, tokenExp int, log *logrus.Logger) *UserController {
	return &UserController{
		repo:            UR,
		pc:              pc,
		secretKey:       secretKey,
		loginLen:        loginLen,
		passwordLen:     passwordLen,
		tokenExpiration: time.Hour * time.Duration(tokenExp),
		log:             log,
	}
}

func (u *UserController) Registration(ctx context.Context, newUser *model.User) (string, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL getWatchingEpisode")
	defer span.End()
	res, err := u.repo.GetUser(ctx, newUser.Username)
	if err != nil {
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("getUser: %w", err)
	}

	if res != nil {
		u.log.Warnf("registation err %s, value %v", errors.ErrorUserExist, newUser)
		return "", fmt.Errorf("%w", errors.ErrorUserExist)
	}

	if len(newUser.Username) < u.loginLen {
		err := errors.LoginLenError{LoginLen: u.loginLen}
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("%w", err)
	}

	if len(newUser.Password) < u.passwordLen {
		err := errors.PasswordLenError{PasswordLen: u.passwordLen}
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("%w", err)
	}

	_, err = mail.ParseAddress(newUser.Email)
	if err != nil {
		u.log.Warnf("registation err %s, value %v", errors.ErrorInvalidEmail, newUser)
		return "", fmt.Errorf("%w", errors.ErrorInvalidEmail)
	}

	newUser.RegData = time.Now()
	newUser.LastActive = time.Now().Add(-time.Hour * 24)
	newUser.LastSubscribe = time.Now()
	newUser.Color = "#000000"
	newUser.Emoji = "2b50"

	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	newUser.Password = string(hash)

	err = u.repo.CreateUser(ctx, newUser)
	if err != nil {
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("createUser: %w", err)
	}

	err = u.pc.EarnPointForLogin(ctx, newUser)
	if err != nil {
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("earnPointForLogin: %w", err)
	}

	newUser.LastActive = time.Now()
	err = u.repo.UpdateUser(ctx, *newUser)
	if err != nil {
		u.log.Warnf("registation err %s, value %v", err, newUser)
		return "", fmt.Errorf("updateUser: %w", err)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(u.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        newUser.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(u.secretKey))

	u.log.Infof("registation user %s", newUser.Username)
	return ss, nil
}

func (u *UserController) Login(ctx context.Context, username, password string) (string, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL Login")
	defer span.End()
	user, err := u.repo.GetUser(ctx, username)
	if err != nil {
		u.log.Warnf("login err %s, value %s", err, username)
		return "", fmt.Errorf("getUser: %w", err)
	}

	if user == nil {
		u.log.Warnf("login err %s, value %s", err, username)
		return "", fmt.Errorf("%w", errors.ErrorWrongLogin)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		u.log.Warnf("login err %s, value %s", err, username)
		return "", fmt.Errorf("%w", errors.ErrorWrongLogin)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(u.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(u.secretKey))

	err = u.UpdateActive(ctx, ss)
	if err != nil {
		u.log.Warnf("login err %s, value %s", err, username)
		return "", fmt.Errorf("updateActite: %w", err)
	}
	u.log.Infof("login user %s", user.Username)
	return ss, nil
}

func eqDate(ctx context.Context, date1, date2 time.Time) bool {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL eqDate")
	defer span.End()
	d1, m1, y1 := date1.Date()
	d2, m2, y2 := date2.Date()

	if d1 != d2 || m1 != m2 || y1 != y2 {
		return false
	}
	return true
}

func (u *UserController) UpdateActive(ctx context.Context, token string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL UpdateActive")
	defer span.End()
	user, err := u.AuthByToken(ctx, token)
	if err != nil {
		u.log.Warnf("update active user auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !eqDate(ctx, user.LastActive, time.Now()) {
		err = u.pc.EarnPointForLogin(ctx, user)
		if err != nil {
			u.log.Warnf("update active user err %s, username %s", err, user.Username)
			return fmt.Errorf("earnPointForLogin: %w", err)
		}
	}

	user.LastActive = time.Now()
	err = u.repo.UpdateUser(ctx, *user)
	if err != nil {
		u.log.Warnf("update active user err %s, username %s", err, user.Username)
		return fmt.Errorf("updateUser: %w", err)
	}
	u.log.Infof("update active user %s", user.Username)
	return nil
}

func (u *UserController) AuthByToken(ctx context.Context, token string) (*model.User, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL AuthByToken")
	defer span.End()
	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.secretKey), nil
	})

	if err != nil {
		u.log.Warnf("auth by token err %s, username %s", err, claims.ID)
		return nil, fmt.Errorf("AuthToken: %w", err)
	}

	user, err := u.repo.GetUser(ctx, claims.ID)
	if err != nil || user == nil {
		u.log.Warnf("auth by token err %s, username %s", err, claims.ID)
		return nil, fmt.Errorf("getUser: %w", err)
	}
	u.log.Infof("auth by token user %s", user.Username)
	return user, nil
}

func (u *UserController) ChangeEmoji(ctx context.Context, token, emojiCode string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL ChangeEmoji")
	defer span.End()
	user, err := u.AuthByToken(ctx, token)
	if err != nil {
		u.log.Warnf("change emoji auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	user.Emoji = emojiCode
	err = u.repo.UpdateUser(ctx, *user)
	if err != nil {
		u.log.Warnf("chande emoji err %s, user %s, value %s", err, user.Username, emojiCode)
		return fmt.Errorf("updateUser: %w", err)
	}
	u.log.Infof("user %s change emoji to %s", user.Username, emojiCode)
	return nil
}

func (u *UserController) ChangeAvatarColor(ctx context.Context, token, color string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL ChangeAvatarColor")
	defer span.End()
	user, err := u.AuthByToken(ctx, token)
	if err != nil {
		u.log.Warnf("change acatar color auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	user.Color = color
	err = u.repo.UpdateUser(ctx, *user)
	if err != nil {
		u.log.Warnf("chande emoji err %s, user %s, value %s", err, user.Username, color)
		return fmt.Errorf("updateUser: %w", err)
	}
	u.log.Infof("user %s change color avatar to %s", user.Username, color)
	return nil
}

func (u *UserController) GetPublicInfo(ctx context.Context, username string) (*model.User, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetPublicInfo")
	defer span.End()
	info, err := u.repo.GetPublicInfo(ctx, username)
	if err != nil {
		u.log.Warnf("get public info err %s about user %s", err, username)
		return nil, fmt.Errorf("getPublicInfo: %w", err)
	}
	return info, nil
}

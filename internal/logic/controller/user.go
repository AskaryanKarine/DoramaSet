package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	repository "DoramaSet/internal/repository/interfaces"
	"time"
)

type UserController struct {
	repo repository.IUserRepo
	pc   interfaces.IPointsController
}

func (u *UserController) Registration(record model.User) error {
	return nil
}
func (u *UserController) Login(username, password string) error {
	return nil
}

func (u *UserController) Logout(username string) error {
	return nil
}

func eqDate(date1, date2 time.Time) bool {
	d1, m1, y1 := date1.Date()
	d2, m2, y2 := date2.Date()

	if d1 != d2 || m1 != m2 || y1 != y2 {
		return false
	}
	return true
}

func (u *UserController) UpdateActive(username string) error {
	user, err := u.repo.GetUser(username)
	if err != nil {
		return err
	}
	if !eqDate(user.LastActive, time.Now()) {
		u.pc.EarnPointForLogin(username)
	}

	user.LastActive = time.Now()
	return u.repo.UpdateUser(user)
}

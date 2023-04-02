package controller

import (
	"DoramaSet/internal/repository/interfaces"
	"errors"
	"time"
)

type PointsController struct {
	repo interfaces.IUserRepo
}

func checkYear(date time.Time) bool {
	today := time.Now()
	if date.Month() != today.Month() {
		return false
	}

	if date.Month() == time.February && date.Day() == 29 {
		date.Add(-time.Hour * 24)
	}

	if date.Day() != today.Day() {
		return false
	}

	return true
}

func (p *PointsController) EarnPointForLogin(username string) error {
	user, err := p.repo.GetUser(username)
	if err != nil {
		return err
	}
	user.Points += 5
	if checkYear(user.RegData) {
		user.Points += 10
	}

	return p.repo.UpdateUser(user)
}

func (p *PointsController) PurgePoint(username string, point int) error {
	user, err := p.repo.GetUser(username)
	if err != nil {
		return err
	}

	if user.Points < point {
		return errors.New("not enough points")
	}

	user.Points -= point
	return p.repo.UpdateUser(user)
}

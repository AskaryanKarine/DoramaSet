package controller

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"time"
)

type PointsController struct {
	repo repository.IUserRepo
}

const (
	everyDayPoint    = 5
	everyYearPoint   = 10
	longNoLoginPoint = 50
	longNoLoginHours = 4400.0
)

func checkYear(date time.Time) bool {
	today := time.Now()

	if date.Month() == time.February && date.Day() == 29 {
		date = date.Add(-time.Hour * 24)
	}

	if date.Day() != today.Day() {
		return false
	}

	if date.Month() != today.Month() {
		return false
	}

	return true
}

func (p *PointsController) EarnPointForLogin(user *model.User) error {
	user.Points += everyDayPoint

	if checkYear(user.RegData) {
		user.Points += everyYearPoint
	}

	if time.Since(user.LastActive).Hours() > longNoLoginHours {
		user.Points += longNoLoginPoint
	}

	err := p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}
	return nil
}

func (p *PointsController) PurgePoint(user *model.User, point int) error {
	if user.Points < point {
		err := errors.BalanceError{
			Have: user.Points,
			Want: point,
		}
		return fmt.Errorf("purgePoint: %w", err)
	}

	user.Points -= point

	err := p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}
	return nil
}

func (p *PointsController) EarnPoint(user *model.User, point int) error {
	user.Points += point
	err := p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}
	return nil
}
